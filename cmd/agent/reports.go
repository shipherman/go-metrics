package main

import (
    "fmt"
    "runtime"
    "strings"
    "net/http"
    "math/rand"
    "encoding/json"
    "bytes"
    "io"

    "github.com/shipherman/go-metrics/internal/storage"
)


type Metrics struct {
    ID    string   `json:"id"`              // имя метрики
    MType string   `json:"type"`            // параметр, принимающий значение gauge или counter
    Delta storage.Counter   `json:"delta"` // значение метрики в случае передачи counter
    Value storage.Gauge `json:"value"` // значение метрики в случае передачи gauge
}

const contentType string = "text/plain"
const counterType string = "counter"
const gaugeType string = "gauge"


func readMemStats(m *storage.MemStorage) {
    var stat runtime.MemStats
    runtime.ReadMemStats(&stat)
    m.UpdateGauge("Alloc", storage.Gauge(stat.Alloc))
    m.UpdateGauge("BuckHashSys", storage.Gauge(stat.BuckHashSys))
    m.UpdateGauge("Frees", storage.Gauge(stat.Frees))
    m.UpdateGauge("GCCPUFraction", storage.Gauge(stat.GCCPUFraction))
    m.UpdateGauge("GCSys", storage.Gauge(stat.GCSys))
    m.UpdateGauge("HeapAlloc", storage.Gauge(stat.HeapAlloc))
    m.UpdateGauge("HeapIdle", storage.Gauge(stat.HeapIdle))
    m.UpdateGauge("HeapInuse", storage.Gauge(stat.HeapInuse))
    m.UpdateGauge("HeapObjects", storage.Gauge(stat.HeapObjects))
    m.UpdateGauge("HeapReleased", storage.Gauge(stat.HeapReleased))
    m.UpdateGauge("HeapSys", storage.Gauge(stat.HeapSys))
    m.UpdateGauge("LastGC", storage.Gauge(stat.LastGC))
    m.UpdateGauge("Lookups", storage.Gauge(stat.Lookups))
    m.UpdateGauge("MCacheInuse", storage.Gauge(stat.MCacheInuse))
    m.UpdateGauge("MCacheSys", storage.Gauge(stat.MCacheSys))
    m.UpdateGauge("MSpanInuse", storage.Gauge(stat.MSpanInuse))
    m.UpdateGauge("MSpanSys", storage.Gauge(stat.MSpanSys))
    m.UpdateGauge("Mallocs", storage.Gauge(stat.Mallocs))
    m.UpdateGauge("NextGC", storage.Gauge(stat.NextGC))
    m.UpdateGauge("NumForcedGC", storage.Gauge(stat.NumForcedGC))
    m.UpdateGauge("NumGC", storage.Gauge(stat.NumGC))
    m.UpdateGauge("OtherSys", storage.Gauge(stat.OtherSys))
    m.UpdateGauge("PauseTotalNs", storage.Gauge(stat.PauseTotalNs))
    m.UpdateGauge("StackInuse", storage.Gauge(stat.StackInuse))
    m.UpdateGauge("StackSys", storage.Gauge(stat.StackSys))
    m.UpdateGauge("Sys", storage.Gauge(stat.Sys))
    m.UpdateGauge("TotalAlloc", storage.Gauge(stat.TotalAlloc))
    m.UpdateGauge("RandomValue", storage.Gauge(rand.Float32()))
    m.UpdateCounter("PollCount", storage.Counter(1))
}

func ProcessReport (serverAddress string, m storage.MemStorage) error {
    // metric type variable

    var metrics Metrics

    serverAddress = strings.Join([]string{"http:/",serverAddress,"update/"}, "/")

    //send request to the server
    for k, v := range m.Data{
        switch v.(type){
            case storage.Gauge:
                metrics = Metrics{ID:k, MType:gaugeType, Value:v.(storage.Gauge)}
            case storage.Counter:
                metrics = Metrics{ID:k, MType:counterType, Delta:v.(storage.Counter)}
            default:
                return fmt.Errorf("uknown type of metric")
        }

        data, err := json.Marshal(metrics)
        if err != nil {
            return err
        }

//         fmt.Println(string(data))

        request, err := http.NewRequest("POST", serverAddress, bytes.NewBuffer(data))
        request.Header.Set("Content-Type", contentType)

        client := &http.Client{}
        resp, err := client.Do(request)

        if err != nil {
            return err
        }

        if resp.StatusCode != http.StatusOK {
            b, _ := io.ReadAll(resp.Body)
            return fmt.Errorf("%s: %s; %s",
                            "Can't send report to the server",
                            resp.Status,
                            b)
        }

        defer resp.Body.Close()

    }
    return nil
}

