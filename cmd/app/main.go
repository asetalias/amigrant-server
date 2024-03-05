package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/asetalias/amigrant-server/config"
	"github.com/asetalias/amigrant-server/internal/app"
	"go.uber.org/zap"
)

func init() {
	log.Println("Initializing ZAP!")
	zap.ReplaceGlobals(zap.Must(zap.NewProduction()))

	zap.L().Info("ZAP is initialized!")
}

func main() {
	zap.L().Info("Initializing config...")
	config := config.LoadConfig()

	zap.L().Info("Initializing server...")
	server := app.NewServer()

	zap.L().Info("Starting server...")
	server.Start(config.Port)

	// Wait for SIGINT (Ctrl+C) or SIGKILL (kill -9 <pid>) to stop the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	zap.L().Info("Server stopping...")

	// close the DB connections, etc.

	os.Exit(0)

}
