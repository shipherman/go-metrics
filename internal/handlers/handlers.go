package handlers

import (
    "fmt"
    "net/http"
    "strconv"

    "github.com/go-chi/chi/v5"
    "github.com/shipherman/go-metrics/internal/storage"
)

type handler struct {
    store storage.MemStorage
}

func NewHandler() handler {
    return handler{store: storage.New()}
}

const counterType = "counter"
const gaugeType = "gauge"


func (h *handler) HandleMain(w http.ResponseWriter, r *http.Request) {
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
    list := h.store.GetAll()
    for k, v := range list {
        body = body + fmt.Sprintf("<tr>\n<td>%s</td>\n", k)
        body = body + fmt.Sprintf("<td>%v</td>\n</tr>\n", v)
    }
    body = body + " </table>\n </body>\n</html>"

    w.Write([]byte(body))
}

func (h *handler) HandleUpdate (w http.ResponseWriter, r *http.Request) {
    // get context params
    metricType := chi.URLParam(r, "type")
    metric := chi.URLParam(r, "metric")
    value := chi.URLParam(r, "value")

    // find out metric type
    switch metricType {
        case counterType:
            v, err := strconv.Atoi(value)
            if err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
            }
            h.store.UpdateCounter(metric, storage.Counter(v))
        case gaugeType:
            v, err := strconv.ParseFloat(value, 64)
            if err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
            }
            h.store.UpdateGauge(metric, storage.Gauge(v))
        default:
            http.Error(w, "Incorrect metric type", http.StatusBadRequest)
    }
}

func (h *handler) HandleValue(w http.ResponseWriter, r *http.Request) {
    metric := chi.URLParam(r, "metric")
    v, err := h.store.Get(metric)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
    }
    fmt.Fprint(w, v)
}

