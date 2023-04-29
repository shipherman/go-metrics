package main

import (
     "fmt"
    "net/http"
//     "io"
//     "encoding/json"
    "strings"
    "strconv"
)


type counter int64
type gauge float64

type MemStorage struct {
    counterData map[string]counter
    gaugeData map[string]gauge
}

var mem = MemStorage{
    map[string]counter{},
    map[string]gauge{},
}


func handleMain(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
}

func handleUpdate (w http.ResponseWriter, r *http.Request) {
    //http://<АДРЕС_СЕРВЕРА>/update/<ТИП_МЕТРИКИ>/<ИМЯ_МЕТРИКИ>/<ЗНАЧЕНИЕ_МЕТРИКИ>,
    //Content-Type: text/plain

    if r.Method != http.MethodPost {
        http.Error(w, "Incorrect HTTP method", http.StatusBadRequest)
    }

    url := strings.Split(r.URL.Path,"/")
    if len(url) < 5 {
        http.Error(w, "Missed data in POST Request", http.StatusBadRequest)
        return
    }
    switch url[2] {
        case "counter":
            i, err := strconv.Atoi(url[4])
            if err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
            } else {
                mem.counterData[url[3]] += counter(i)
                w.WriteHeader(http.StatusOK)
            }
        case "gauge":
            i, err := strconv.ParseFloat(url[4], 64)
            if err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
            } else {
                mem.gaugeData[url[3]] = gauge(i)
                w.WriteHeader(http.StatusOK)
            }
        default:
            http.Error(w, "Incorrect metric type", http.StatusBadRequest)
    }
    fmt.Println(&mem)
}

func main() {
    //run server
    server := http.NewServeMux()
    server.HandleFunc(`/`, handleMain)
    server.HandleFunc(`/update/`, handleUpdate)
    err := http.ListenAndServe(":8080", server)
    if err != nil {
        panic(err)
    }
}
