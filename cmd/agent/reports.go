package main

import (
    "fmt"
    "io"
    "runtime"
    "strings"
    "net/http"
    "math/rand"

    "github.com/shipherman/go-metrics/internal/storage"
)

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

func sendReport (req string) error {
    reader := new(io.Reader)
    resp, err := http.Post(req, contentType, *reader)
    if err != nil {
        return err
    }
    if resp.StatusCode != http.StatusOK {
        line, err := io.ReadAll(resp.Body)
        if err != nil {
            return err
        }
        fmt.Println(req)
        return fmt.Errorf("%s: %s; %s",
                          "Can't send report to the server",
                          resp.Status,
                          line)
    }
    resp.Body.Close()
    return nil
}

func ProcessReport (serverAddress string, m storage.MemStorage) error {
    // metric type variable
    var mtype string

    //send request to the server
    for k, v := range m.Data{
        switch v.(type){
            case storage.Gauge:
                mtype = gaugeType //replace with const string
            case storage.Counter:
                mtype = counterType
            default:
                return fmt.Errorf("uknown type of metric")
        }
        req := strings.Join([]string{"http:/",
                         serverAddress,
                         "update",
                         mtype,
                         fmt.Sprintf("%v/%v",k,v)}, "/")
        err := sendReport(req)
        if err != nil {
            return err
        }
    }
    return nil
}

