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
    "log"

    "compress/gzip"

    "github.com/shipherman/go-metrics/internal/storage"
)


type Metrics struct {
    ID    string   `json:"id"`              // имя метрики
    MType string   `json:"type"`            // параметр, принимающий значение gauge или counter
    Delta storage.Counter   `json:"delta"` // значение метрики в случае передачи counter
    Value storage.Gauge `json:"value"` // значение метрики в случае передачи gauge
}

const contentType string = "application/json"
const compression string = "gzip"

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

func Compress(data []byte) ([]byte, error) {
    var b bytes.Buffer
    w, err := gzip.NewWriterLevel(&b, gzip.BestSpeed)
    if err != nil {
        return nil, fmt.Errorf("failed init compress writer: %v", err)
    }
    _, err = w.Write(data)
    if err != nil {
        return nil, fmt.Errorf("failed write data to compress temporary buffer: %v", err)
    }
    err = w.Close()
    if err != nil {
        return nil, fmt.Errorf("failed compress data: %v", err)
    }
    return b.Bytes(), nil
}

func sendReport (serverAddress string, metrics Metrics) error {
    data, err := json.Marshal(metrics)
    if err != nil {
        return err
    }

    data, err = Compress(data)
    if err != nil {
        return err
    }

    request, err := http.NewRequest("POST", serverAddress, bytes.NewBuffer(data))
    if err != nil {
        return err
    }
    request.Header.Set("Content-Type", contentType)
    request.Header.Set("Content-Encoding", compression)
    request.Header.Set("Accept-Encoding", compression)

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
    return nil
}

func ProcessReport (serverAddress string, m storage.MemStorage) error {
    // metric type variable

    var metrics Metrics

    serverAddress = strings.Join([]string{"http:/",serverAddress,"update/"}, "/")

    //send request to the server
    for k, v := range m.CounterData{
        metrics = Metrics{ID:k, MType:counterType, Delta:v}
        log.Println(metrics)
        err := sendReport(serverAddress, metrics)
        if err != nil {
            return err
        }
    }

    for k, v := range m.GaugeData{
        metrics = Metrics{ID:k, MType:gaugeType, Value:v}
        err := sendReport(serverAddress, metrics)
        if err != nil {
            return err
        }
    }

    return nil
}

