package main

import (
    "io"
    "fmt"
    "net/http"
    "log"
    "strconv"

    "github.com/go-chi/chi/v5"
    s "github.com/shipherman/go-metrics/internal/storage"
)


var mem = s.MemStorage{
    Data: map[string]interface{}{},
}


func HandleMain(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "BadRequest", http.StatusBadRequest)
}

func HandleUpdate (w http.ResponseWriter, r *http.Request) {
    //http://<АДРЕС_СЕРВЕРА>/update/<ТИП_МЕТРИКИ>/<ИМЯ_МЕТРИКИ>/<ЗНАЧЕНИЕ_МЕТРИКИ>,
    //Content-Type: text/plain

    metricType := chi.URLParam(r, "type")
    metric := chi.URLParam(r, "metric")
    value := chi.URLParam(r, "value")

    // find out metric type
    switch metricType {
        case "counter":
            v, err := strconv.Atoi(value)
            if err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
            }
            mem.Update(metricType, metric, s.Counter(v))
        case "gauge":
            v, err := strconv.ParseFloat(value, 64)
            if err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
            }
            mem.Update(metricType, metric, s.Gauge(v))
        default:
            http.Error(w, "Incorrect metric type", http.StatusBadRequest)
    }
//      fmt.Println(&mem)
}

func HandleValue (w http.ResponseWriter, r *http.Request) {
    metric := chi.URLParam(r, "metric")
    v, err := mem.Get(metric)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
    }
    io.WriteString(w, fmt.Sprintf("%v", v))
}

func main() {
    // Routers
    router := chi.NewRouter()
    router.Get("/", HandleMain)
    router.Post("/update/{type}/{metric}/{value}", HandleUpdate)
    router.Get("/value/gauge/{metric}", HandleValue)
    router.Get("/value/counter/{metric}", HandleValue)

    //run server
    log.Fatal(http.ListenAndServe(":8080", router))
}
