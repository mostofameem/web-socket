package handlers

import (
	"net/http"
	"websocket/rest/utils"
)

func (h *Handlers) Hello(w http.ResponseWriter, r *http.Request) {
	utils.SendJson(w, http.StatusOK, map[string]any{
		"success": true,
	})
}
