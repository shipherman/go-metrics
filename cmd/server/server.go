package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/shipherman/go-metrics/internal/db"
	"github.com/shipherman/go-metrics/internal/handlers"
	"github.com/shipherman/go-metrics/internal/options"
	"github.com/shipherman/go-metrics/internal/routers"
	"github.com/shipherman/go-metrics/internal/storage"
)

var buildVersion string
var buildDate string
var buildCommit string

func main() {
	log.Println("Starting server...")
	fmt.Printf("version=%s, builddate=%s, commit=%s\n", buildVersion, buildDate, buildCommit)
	// Store variable will be used file or database to save metrics
	var store storage.StorageWriter

	// Parse cli options into config
	cfg, err := options.ParseOptions()
	if err != nil {
		panic(err)
	}

	log.Println("Params:", cfg)

	// Handler for router
	h := handlers.NewHandler()

	// Identify wether use DB or file to save metrics
	if cfg.DBDSN != "" {
		database, err := db.Connect(cfg.DBDSN)
		if err != nil {
			log.Println(err)
		}

		// Use database as a store
		store = &database

		//Define DB for handlers
		h.DBconn = database.Conn

	} else {
		// use json file to store metrics
		store = &storage.Localfile{Path: cfg.Filename}
	}

	// Init router
	router, err := routers.InitRouter(cfg, h)
	if err != nil {
		panic(err)
	}

	if cfg.Restore {
		err := store.RestoreData(&h.Store)
		if err != nil {
			log.Println("Could not restore data: ", err)
		}
	}

	// Write MemStorage to a store provider
	// Interval used for file saving
	go func() {
		for {
			store.Save(cfg.Interval, h.Store)
		}
	}()

	// Define server parameters
	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	log.Println("Started. Running")

	// Graceful shutdown
	idleConnectionsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-sigint
		log.Println("Shutting down server")

		if err := store.Write(h.Store); err != nil {
			log.Printf("Error during saving data to file: %v", err)
		}

		// Close file/db
		defer store.Close()

		if err := server.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP Server Shutdown Error: %v", err)
		}
		close(idleConnectionsClosed)
	}()

	// Run server
	log.Fatal(server.ListenAndServe())

	<-idleConnectionsClosed
	log.Println("Server shutdown")
}
