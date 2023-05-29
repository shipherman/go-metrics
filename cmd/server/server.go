package server

import (
    "os"
    "os/signal"
    "syscall"
    "net/http"
    "log"
    "context"

    "github.com/shipherman/go-metrics/internal/routers"
    "github.com/shipherman/go-metrics/internal/options"
)


func main() {
    //parse cli options
    cfg, err := options.ParseOptions()
    if err != nil {
        panic(err)
    }

    log.Println(cfg)
    log.Println("Starting server...")


    router, hStore, err := routers.InitRouter(cfg)
    if err != nil {
        panic(err)
    }

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

        if err := hStore.SaveDataToFile(); err != nil {
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
