package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/keithyw/pitch-in/internal/config"
	"github.com/keithyw/pitch-in/internal/database"
	"github.com/keithyw/pitch-in/internal/server"
	"github.com/keithyw/pitch-in/pkg/storage"
)

func main() {
	ctx := context.Background()
	log := slog.Default()
	config := config.NewConfig()

	var client database.DBClient
	var err error

	maxRetries := 10

	for i := 0; i < maxRetries; i++ {
		client, err = database.NewDBClient(config)
		if err == nil {
			err = client.Ping()
		}

		if err == nil {
			log.Info("Successfully connected to Mysql")
			break
		}

		log.Warn("MySQL not ready, retrying...", "attempt", i+1, "error", err)
		time.Sleep(3 * time.Second)

		if i == maxRetries-1 {
			panic("Failed loading mysql after multiple retries: " + err.Error())
		}
	}

	var storageClient storage.StorageClient
	storageConfig := storage.NewStorageConfig()

	for i := 0; i < maxRetries; i++ {
		storageClient, err = storage.NewMinioClient(ctx, storageConfig)
		if err == nil {
			log.Info("Successfully connected to Storage Client")
			break
		}
		log.Warn("Storage Client not ready, retrying...", "attempt", i+1, "error", err)
		time.Sleep(3 * time.Second)
		if i == maxRetries-1 {
			panic("Failed loading minio storage client: " + err.Error())
		}
	}

	store := database.NewDBStore(ctx, client)
	server := server.NewServer(config, store, storageClient, log)
	log.Info(fmt.Sprintf("Server starting on %s", config.HttpPort))
	http.ListenAndServe(config.HttpPort, server)
}
