package rest

import (
	"net/http"

	"websocket/rest/middlewares"
)

func (s *Server) initRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle(
		"GET /websocket/api/v1/hello",
		manager.With(
			http.HandlerFunc(s.handlers.Hello),
		),
	)
	mux.HandleFunc(
		"/websocket/hello",
		s.handlers.WsHandler,
	)

}
