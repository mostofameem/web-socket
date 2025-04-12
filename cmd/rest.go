package cmd

import (
	"log/slog"

	"websocket/config"
	"websocket/logger"
	"websocket/rest"
	"websocket/rest/handlers"
	"websocket/rest/utils"

	"github.com/spf13/cobra"
)

var serveRestCmd = &cobra.Command{
	Use:   "serve-rest",
	Short: "Start the websocket service",
	RunE:  serveRest,
}

func serveRest(cmd *cobra.Command, args []string) error {
	conf := config.GetConfig()

	//apm.InitAPMClient(*conf.Apm)

	logger.SetupLogger(conf.ServiceName)
	utils.InitValidator()


	slog.Info("Redis client is connected.")

	handler := handlers.NewHandler(conf/*, agentSvc*/)

	server := rest.NewServer(conf, handler)

	server.Start()
	server.Wg.Wait()

	return nil
}
