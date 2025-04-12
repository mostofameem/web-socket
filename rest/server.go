package rest

import (
	"fmt"
	"log/slog"
	"net/http"
	"sync"
	"websocket/config"
	"websocket/rest/handlers"
	"websocket/rest/middlewares"

	"go.elastic.co/apm/module/apmhttp"
)

type Server struct {
	handlers *handlers.Handlers
	conf     *config.Config
	Wg       sync.WaitGroup
}

func NewServer(conf *config.Config, handlers *handlers.Handlers) *Server {
	return &Server{
		conf:     conf,
		handlers: handlers,
	}
}

func (s *Server) Start() {

	manager := middlewares.NewManager()

	manager.Use(
		middlewares.Recover,
		middlewares.Logger,
	)

	mux := http.NewServeMux()

	s.initRoutes(mux, manager)

	handler := middlewares.EnableCors(mux)

	// swagger.SetupSwagger(s.conf, mux, manager)

	s.Wg.Add(1)

	go func() {
		defer s.Wg.Done()

		addr := fmt.Sprintf(":%d", s.conf.HttpPort)
		slog.Info(fmt.Sprintf("Listening at %s", addr))

		if err := http.ListenAndServe(addr, apmhttp.Wrap(handler)); err != nil {
			slog.Error(err.Error())
		}
	}()
}
