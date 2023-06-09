package handlers

import (
    "fmt"
    "net/http"
    "context"
)


func (h *Handler) HandlePing(w http.ResponseWriter, r *http.Request) {
    v := "pong\n"
    err := dbconn.Ping(context.Background())
    if err != nil {
        http.Error(w, "Connection to DB is lost", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, v)
}
