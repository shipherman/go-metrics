package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

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
		ctx := context.Background()
		database, err := db.Connect(ctx, cfg.DBDSN)
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
			err = store.Save(cfg.Interval, h.Store)
			if err != nil {
				log.Println("Could not save data: ", err)
			}
		}
	}()

	// Define server parameters
	server := http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		WriteTimeout: time.Second * 10,
	}

	// Register func to save data on Shutdown
	// Add WaitGroup to sync shutdown
	var wg sync.WaitGroup
	wg.Add(1)
	server.RegisterOnShutdown(func() {
		if err := store.Write(h.Store); err != nil {
			log.Printf("Error during saving data to file: %v", err)
		}
		wg.Done()
	})

	log.Println("Started. Running")

	// Graceful shutdown here
	idleConnectionsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		<-sigint
		log.Println("Shutting down server")

		// Shutdown server
		if err := server.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP Server Shutdown Error: %v", err)
		}

		// Wait till registred shutdown function will be done
		wg.Wait()

		// Close file/db
		store.Close()

		close(idleConnectionsClosed)
		log.Printf("HTTP server shutted down")
	}()

	// Run server
	if err := server.ListenAndServe(); err != nil {
		if err.Error() != "http: Server closed" {
			log.Printf("HTTP server closed with: %v\n", err)
		}
	}
	<-idleConnectionsClosed
}
