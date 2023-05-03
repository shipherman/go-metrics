package main

import (
//      "fmt"
    "net/http"
//     "io"
//     "encoding/json"
    "strings"
    "strconv"
    s "github.com/shipherman/go-metrics/internal/storage"
)


var mem = s.MemStorage{
    CounterData: map[string]s.Counter{},
    GaugeData: map[string]s.Gauge{},
}


func HandleMain(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "BadRequest", http.StatusBadRequest)
}

func HandleUpdate (w http.ResponseWriter, r *http.Request) {
    //http://<АДРЕС_СЕРВЕРА>/update/<ТИП_МЕТРИКИ>/<ИМЯ_МЕТРИКИ>/<ЗНАЧЕНИЕ_МЕТРИКИ>,
    //Content-Type: text/plain

    if r.Method != http.MethodPost {
        http.Error(w, "Incorrect HTTP method", http.StatusMethodNotAllowed)
    }

    url := strings.Split(r.URL.Path,"/")
    if len(url) < 5 {
        http.Error(w, "Missed data in POST Request", http.StatusNotFound)
        return
    }
    switch url[2] {
        case "counter":
            i, err := strconv.Atoi(url[4])
            if err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
            } else {
                mem.CounterData[url[3]] += s.Counter(i)
                w.WriteHeader(http.StatusOK)
            }
        case "gauge":
            i, err := strconv.ParseFloat(url[4], 64)
            if err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
            } else {
                mem.GaugeData[url[3]] = s.Gauge(i)
                w.WriteHeader(http.StatusOK)
            }
        default:
            http.Error(w, "Incorrect metric type", http.StatusBadRequest)
    }
//      fmt.Println(&mem)
}

func main() {
    //run server
    server := http.NewServeMux()
    server.HandleFunc(`/`, HandleMain)
    server.HandleFunc(`/update/`, HandleUpdate)
    err := http.ListenAndServe(":8080", server)
    if err != nil {
        panic(err)
    }
}
