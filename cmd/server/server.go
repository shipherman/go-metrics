package main

import (
    "io"
    "os"
    "fmt"
    "flag"
    "net/http"
    "log"
    "strconv"

    "github.com/go-chi/chi/v5"
    s "github.com/shipherman/go-metrics/internal/storage"
)

//storage
var mem = s.MemStorage{
    Data: map[string]interface{}{},
}

//cli options
var options struct {
    address string
}


func HandleMain(w http.ResponseWriter, r *http.Request) {
    //write static html page with all the items to the response; unsorted
    body := `
<!DOCTYPE html>
<html>
    <head>
        <title>All tuples</title>
    </head>
    <body>
      <table>
          <tr>
            <td>Metric</td>
            <td>Value</td>
          </tr>
    `
    list := mem.GetAll()
    for k, v := range list {
        body = body + fmt.Sprintf("<tr>\n<td>%s</td>\n", k)
        body = body + fmt.Sprintf("<td>%v</td>\n</tr>\n", v)
    }
    body = body + " </table>\n </body>\n</html>"

    w.Write([]byte(body))
}

func HandleUpdate (w http.ResponseWriter, r *http.Request) {
    //http://<АДРЕС_СЕРВЕРА>/update/<ТИП_МЕТРИКИ>/<ИМЯ_МЕТРИКИ>/<ЗНАЧЕНИЕ_МЕТРИКИ>,
    //Content-Type: text/plain

    // get context params
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
    router.Get("/", HandleMain)
    router.Post("/update/{type}/{metric}/{value}", HandleUpdate)
    router.Get("/value/gauge/{metric}", HandleValue)
    router.Get("/value/counter/{metric}", HandleValue)

    log.Println("Starting server...")
    //run server
    log.Fatal(http.ListenAndServe(options.address, router))
}
