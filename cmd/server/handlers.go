package main

import (
    "fmt"
    "net/http"
    "strconv"

    "github.com/go-chi/chi/v5"
    s "github.com/shipherman/go-metrics/internal/storage"
)

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

func HandleValue(w http.ResponseWriter, r *http.Request) {
    metric := chi.URLParam(r, "metric")
    v, err := mem.Get(metric)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
    }
    fmt.Fprint(w, v)
}
