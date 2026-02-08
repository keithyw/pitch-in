package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/keithyw/pitch-in/internal/config"
	"github.com/keithyw/pitch-in/internal/database"
	"github.com/keithyw/pitch-in/internal/server"
)

func main() {
	ctx := context.Background()
	log := slog.Default()
	config := config.NewConfig()
	client, err := database.NewDBClient(config)
	if err != nil {
		panic("Failed loading mysql: " + err.Error())
	}

	store := database.NewDBStore(ctx, client)
	server := server.NewServer(store, log)
	log.Info(fmt.Sprintf("Server starting on %s", config.HttpPort))
	http.ListenAndServe(config.HttpPort, server)
}
