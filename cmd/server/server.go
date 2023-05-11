package main

import (
    "net/http"
    "log"
)


func main() {
    //parse cli options
    cfg := parseOptions()

    log.Println("Starting server...")

    //run server
    log.Fatal(http.ListenAndServe(cfg.address, initRouter()))
}
