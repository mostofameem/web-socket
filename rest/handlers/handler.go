package handlers

import (
	"websocket/config"
)

type Handlers struct {
	conf *config.Config
	//svc  *agent.Service
}

func NewHandler(conf *config.Config /*, service *agent.Service*/) *Handlers {
	return &Handlers{
		conf: conf,
		//svc:  service,
	}
}
