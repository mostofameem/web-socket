package handlers

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/coder/websocket"
)

type TicketLock struct {
	UserID    string
	Timestamp time.Time
	Cancel    context.CancelFunc
}

var ticketLocks = make(map[string]*TicketLock)
var lockMu sync.Mutex

func (handler *Handlers) WsHandler(w http.ResponseWriter, r *http.Request) {
	ticketID := r.URL.Query().Get("ticketId")
	userID := r.URL.Query().Get("userId")
	if ticketID == "" || userID == "" {
		http.Error(w, "Missing ticketId or userId", http.StatusBadRequest)
		return
	}

	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{InsecureSkipVerify: true})
	if err != nil {
		log.Println("WebSocket accept error:", err)
		return
	}

	ctx, cancel := context.WithCancel(r.Context())

	if !tryLockTicket(ticketID, userID, cancel) {
		conn.Write(ctx, websocket.MessageText, []byte("Ticket is already being viewed by someone else"))
		conn.Close(websocket.StatusPolicyViolation, "Ticket locked")
		return
	}

	defer func() {
		unlockTicket(ticketID, userID)
		cancel()
	}()

	ctx, cancel = context.WithTimeout(r.Context(), 30*time.Second)
	defer cancel()

	go keepAlive(ctx, conn, cancel)

	// Handle incoming messages
	for {
		_, msg, err := conn.Read(ctx)
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		log.Println("Message from", userID, "on ticket", ticketID, ":", string(msg))

		if string(msg) == "cancel" {
			log.Println("User canceled viewing ticket")
			break
		}
	}
}

func keepAlive(ctx context.Context, conn *websocket.Conn, cancel context.CancelFunc) {
	ticker := time.NewTicker(5 * time.Second) // ping every 5 seconds now
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			pingCtx, pingCancel := context.WithTimeout(ctx, 3*time.Second)

			log.Println("Sending ping to client...")

			if err := conn.Ping(pingCtx); err != nil {
				log.Println("Ping error, client may be disconnected:", err)
				cancel()
				pingCancel()
				return
			}

			log.Println("Ping successful, client is still connected.")
			pingCancel()

		case <-ctx.Done():
			log.Println("Context canceled, stopping keepAlive.")
			return
		}
	}
}

func tryLockTicket(ticketID, userID string, cancel context.CancelFunc) bool {
	lockMu.Lock()
	defer lockMu.Unlock()

	// If ticket is already locked by someone else
	if lock, ok := ticketLocks[ticketID]; ok && lock.UserID != userID {
		return false
	}

	ticketLocks[ticketID] = &TicketLock{
		UserID:    userID,
		Timestamp: time.Now(),
		Cancel:    cancel,
	}
	return true
}

func unlockTicket(ticketID, userID string) {
	lockMu.Lock()
	defer lockMu.Unlock()

	if lock, ok := ticketLocks[ticketID]; ok && lock.UserID == userID {
		delete(ticketLocks, ticketID)
		log.Println("Ticket", ticketID, "unlocked by", userID)
	}
}
