package main

import (
    "os"
    "os/signal"
    "syscall"
    "net/http"
    "log"
    "context"
    "time"

    "github.com/shipherman/go-metrics/internal/routers"
    "github.com/shipherman/go-metrics/internal/options"
    "github.com/shipherman/go-metrics/internal/storage"
    "github.com/shipherman/go-metrics/internal/handlers"
    "github.com/shipherman/go-metrics/internal/db"

)


func main() {
    //parse cli options
    cfg, err := options.ParseOptions()
    if err != nil {
        panic(err)
    }

    conn, err := db.Connect(cfg.DBDSN)
    if err != nil {
        log.Println("Could not connect to DB. ERROR: ", err)
    }
    defer conn.Close(context.Background())

    handlers.SetDB(conn)

    log.Println(cfg)
    log.Println("Starting server...")


    router, hStore, err := routers.InitRouter(cfg)
    if err != nil {
        panic(err)
    }

    go func() {
        for {
            time.Sleep(time.Second * time.Duration(cfg.Interval))
            storage.WriteDataToFile(cfg.Filename, hStore.Store)
        }
    }()

    server := http.Server{
        Addr: cfg.Address,
        Handler: router,
    }

    idleConnectionsClosed := make(chan struct{})
    go func() {
        sigint := make(chan os.Signal, 1)
        signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
        <-sigint
        log.Println("Shutting down server")

        if err := storage.WriteDataToFile(cfg.Filename, hStore.Store); err != nil {
            log.Printf("Error during saving data to file: %v", err)
        }

        if err := server.Shutdown(context.Background()); err != nil {
            log.Printf("HTTP Server Shutdown Error: %v", err)
        }
        close(idleConnectionsClosed)
    }()

    //run server
    log.Fatal(server.ListenAndServe())

    <-idleConnectionsClosed
    log.Println("Server shutdown")
}
