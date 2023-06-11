package main

import (
    "os"
    "os/signal"
    "syscall"
    "net/http"
    "log"
    "context"
//     "time"

    "github.com/shipherman/go-metrics/internal/routers"
    "github.com/shipherman/go-metrics/internal/options"
    "github.com/shipherman/go-metrics/internal/storage"
    "github.com/shipherman/go-metrics/internal/handlers"
    "github.com/shipherman/go-metrics/internal/db"

)



func main() {
    // Store variable will be used file or database to save metrics
    var store storage.StorageWriter

    // Parse cli options into config
    cfg, err := options.ParseOptions()
    if err != nil {
        panic(err)
    }

    log.Println("Starting server...")
    log.Println("Params:", cfg)

    // Handler for router
    h := handlers.NewHandler()

    router, err := routers.InitRouter(cfg, h)
    if err != nil {
        panic(err)
    }

    // Identify wether use DB or file to save metrics
    if cfg.DBDSN != "" {
        database, err := db.Connect(cfg.DBDSN)
        if err != nil {
            log.Println(err)
        }

        // Use database as a store
        store = &database

        //Define DB for handlers
        handlers.SetDB(database.Conn)

    } else if cfg.Filename != "" {
        // use json file to store metrics
        store = &storage.Localfile{Path: cfg.Filename}
    }

    // Write MemStorage to a store provider
    // Interval used for file saving
    go func() {
        for {
            store.Save(cfg.Interval, h.Store)
        }
    }()

    // Close file/db
    defer store.Close()

    // Define server parameters
    server := http.Server{
        Addr: cfg.Address,
        Handler: router,
    }

    // Graceful shutdown
    idleConnectionsClosed := make(chan struct{})
    go func() {
        sigint := make(chan os.Signal, 1)
        signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
        <-sigint
        log.Println("Shutting down server")

        if err := storage.SaveData(h.Store, store); err != nil {
            log.Printf("Error during saving data to file: %v", err)
        }

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
