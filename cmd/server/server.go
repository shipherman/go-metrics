package main

import (
    "os"
//     "fmt"
    "flag"
    "net/http"
    "log"

    "github.com/go-chi/chi/v5"

    s "github.com/shipherman/go-metrics/internal/storage"
    h "github.com/shipherman/go-metrics/internal/handlers"
)

//storage
var mem = s.MemStorage{
    Data: map[string]interface{}{},
}

//cli options
var options struct {
    address string
}


func main() {
    //parse cli options
    flag.StringVar(&options.address,
                   "a", "localhost:8080",
                   "Add addres and port in format <address>:<port>")
    flag.Parse()

    // get env vars
    if a := os.Getenv("ADDRESS"); a != "" {
        options.address = a
    }

    // Routers
    router := chi.NewRouter()
    router.Get("/", h.HandleMain(&mem))
    router.Post("/update/{type}/{metric}/{value}", h.HandleUpdate(&mem))
    router.Get("/value/gauge/{metric}", h.HandleValue(&mem))
    router.Get("/value/counter/{metric}", h.HandleValue(&mem))

    log.Println("Starting server...")
    //run server
    log.Fatal(http.ListenAndServe(options.address, router))
}
