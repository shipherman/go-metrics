package handlers

import (
    "fmt"
    "net/http"
    "strconv"
    "encoding/json"
    "bytes"
    "os"
    "io"
    "time"
    "log"

    "github.com/go-chi/chi/v5"

    "github.com/shipherman/go-metrics/internal/storage"
    "github.com/shipherman/go-metrics/internal/storage/filestore"
    "github.com/shipherman/go-metrics/internal/options"
)

type Metrics struct {
    ID    string   `json:"id"`              // имя метрики
    MType string   `json:"type"`            // параметр, принимающий значение gauge или counter
    Delta *int64   `json:"delta,omitempty"` // значение метрики в случае передачи counter
    Value *float64 `json:"value,omitempty"` // значение метрики в случае передачи gauge
}

type Handler struct {
    Store storage.MemStorage
    filename string
    ticker *time.Ticker
}

const counterType = "counter"
const gaugeType = "gauge"


func NewHandler(cfg options.Options) (Handler, error) {
    var h Handler
    h.Store = storage.New()
    h.filename = cfg.Filename
    h.ticker = time.NewTicker(time.Duration(cfg.Interval) * time.Second)

    // Read saved metrics from file
    if cfg.Restore {
        f, err := os.OpenFile(h.filename, os.O_RDONLY | os.O_CREATE, 0666)
        defer f.Close()

        if err != nil {
            return h, err
        }

        data, err := io.ReadAll(f)
        if err != nil {
            return h, err
        }

        err = json.Unmarshal(data, &h.Store)
        if err != nil {
            log.Println("Could not restore data", err)
        }
    }
    return h, nil
}

func (h *Handler) SaveDataToFile() error {
    err := filestore.WriteDataToFile(h.filename, h.Store)
    if err != nil {
        return err
    }
    return nil
}

func (h *Handler) SaveDataToFileOnTimer() error {
    skip := make(chan bool)
    go func(){
        time.Sleep(time.Second)
        skip <- true
    }()

    select {
        case <-skip:
            log.Println("Skip saving")
            return nil
        case  <-h.ticker.C:
            log.Println("Save data to file")
            err := filestore.WriteDataToFile(h.filename, h.Store)
            if err != nil {
                return err
            }
    }
    return nil
}

func (h *Handler) HandleMain(w http.ResponseWriter, r *http.Request) {
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
    listC := h.Store.GetAllCounters()
    for k, v := range listC {
        body = body + fmt.Sprintf("<tr>\n<td>%s</td>\n", k)
        body = body + fmt.Sprintf("<td>%v</td>\n</tr>\n", v)
    }

    listG := h.Store.GetAllGauge()
    for k, v := range listG {
        body = body + fmt.Sprintf("<tr>\n<td>%s</td>\n", k)
        body = body + fmt.Sprintf("<td>%v</td>\n</tr>\n", v)
    }

    body = body + " </table>\n </body>\n</html>"

    // respond to agent
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(body))
}

func (h *Handler) HandleValue(w http.ResponseWriter, r *http.Request) {
    metric := chi.URLParam(r, "metric")
    v, err := h.Store.Get(metric)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
    }
    fmt.Fprint(w, v)
}

func (h *Handler) HandleJSONValue(w http.ResponseWriter, r *http.Request) {
    var m Metrics
    var buf bytes.Buffer

    _, err := buf.ReadFrom(r.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err = json.Unmarshal(buf.Bytes(), &m)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

//     fmt.Println(m.ID, m.MType, *m.Delta, *m.Value)

    switch m.MType {
        case counterType:
            v, ok := h.Store.CounterData[m.ID]
            if !ok {
                http.Error(w, "not found", http.StatusNotFound)
                return
            }
            vPtr := int64(v)
            m.Delta = &vPtr
        case gaugeType:
            v, ok := h.Store.GaugeData[m.ID]
            if !ok {
                http.Error(w, "not found", http.StatusNotFound)
                return
            }
            vPtr := float64(v)
            m.Value = &vPtr
    }


    resp, err := json.Marshal(m)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // respond to agent
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(resp)
}

//Update
func (h *Handler) HandleUpdate(w http.ResponseWriter, r *http.Request) {
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
            h.Store.UpdateCounter(metric, storage.Counter(v))
        case gaugeType:
            v, err := strconv.ParseFloat(value, 64)
            if err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
            }
            h.Store.UpdateGauge(metric, storage.Gauge(v))
        default:
            http.Error(w, "Incorrect metric type", http.StatusBadRequest)
    }
    h.SaveDataToFileOnTimer()
}

func (h *Handler) HandleJSONUpdate(w http.ResponseWriter, r *http.Request) {
    var m Metrics
    var buf bytes.Buffer

    _, err := buf.ReadFrom(r.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

//     fmt.Println("inside HandleJSONUpdate", string(buf.Bytes()))

    err = json.Unmarshal(buf.Bytes(), &m)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

//     fmt.Println(m.ID, m.MType, *m.Delta, *m.Value)

    switch m.MType {
        case counterType:
            if m.Delta == nil {
                http.Error(w, "metric value should not be empty", http.StatusBadRequest)
                return
            }
            h.Store.UpdateCounter(m.ID, storage.Counter(*m.Delta))
            w.WriteHeader(http.StatusOK)
        case gaugeType:
            if m.Value == nil {
                http.Error(w, "metric value should not be empty", http.StatusBadRequest)
                return
            }
            h.Store.UpdateGauge(m.ID, storage.Gauge(*m.Value))
            w.WriteHeader(http.StatusOK)
        default:
            http.Error(w, "Incorrect metric type", http.StatusBadRequest)
    }
    h.SaveDataToFileOnTimer()
}
