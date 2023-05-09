package main

import (
    "fmt"
    "io"
    "runtime"
    "strings"
    "net/http"
    "math/rand"

    s "github.com/shipherman/go-metrics/internal/storage"
)

//init MemStorage
var m = s.MemStorage{Data: make(map[string]interface{})}

//MemStats instance
var stat runtime.MemStats


func readMemStats(m *s.MemStorage) {
    runtime.ReadMemStats(&stat)
    m.Update("gauge", "Alloc", s.Gauge(stat.Alloc))
    m.Update("gauge", "BuckHashSys", s.Gauge(stat.BuckHashSys))
    m.Update("gauge", "Frees", s.Gauge(stat.Frees))
    m.Update("gauge", "GCCPUFraction", s.Gauge(stat.GCCPUFraction))
    m.Update("gauge", "GCSys", s.Gauge(stat.GCSys))
    m.Update("gauge", "HeapAlloc", s.Gauge(stat.HeapAlloc))
    m.Update("gauge", "HeapIdle", s.Gauge(stat.HeapIdle))
    m.Update("gauge", "HeapInuse", s.Gauge(stat.HeapInuse))
    m.Update("gauge", "HeapObjects", s.Gauge(stat.HeapObjects))
    m.Update("gauge", "HeapReleased", s.Gauge(stat.HeapReleased))
    m.Update("gauge", "HeapSys", s.Gauge(stat.HeapSys))
    m.Update("gauge", "LastGC", s.Gauge(stat.LastGC))
    m.Update("gauge", "Lookups", s.Gauge(stat.Lookups))
    m.Update("gauge", "MCacheInuse", s.Gauge(stat.MCacheInuse))
    m.Update("gauge", "MCacheSys", s.Gauge(stat.MCacheSys))
    m.Update("gauge", "MSpanInuse", s.Gauge(stat.MSpanInuse))
    m.Update("gauge", "MSpanSys", s.Gauge(stat.MSpanSys))
    m.Update("gauge", "Mallocs", s.Gauge(stat.Mallocs))
    m.Update("gauge", "NextGC", s.Gauge(stat.NextGC))
    m.Update("gauge", "NumForcedGC", s.Gauge(stat.NumForcedGC))
    m.Update("gauge", "NumGC", s.Gauge(stat.NumGC))
    m.Update("gauge", "OtherSys", s.Gauge(stat.OtherSys))
    m.Update("gauge", "PauseTotalNs", s.Gauge(stat.PauseTotalNs))
    m.Update("gauge", "StackInuse", s.Gauge(stat.StackInuse))
    m.Update("gauge", "StackSys", s.Gauge(stat.StackSys))
    m.Update("gauge", "Sys", s.Gauge(stat.Sys))
    m.Update("gauge", "TotalAlloc", s.Gauge(stat.TotalAlloc))
    m.Update("gauge", "RandomValue", s.Gauge(rand.Float32()))
    m.Update("counter", "PollCount", s.Counter(1))
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
        return fmt.Errorf("%s: %s; %s",
                          "Can't send report to the server",
                          resp.Status,
                          line)
    }
    resp.Body.Close()
    return nil
}

func ProcessReport (data *s.MemStorage) error {
    // metric type variable
    var mtype string

    //send request to the server
    for k, v := range data.Data {
        switch v.(type){
            case s.Gauge:
                mtype = "gauge"
            case s.Counter:
                mtype = "counter"
        }
        req := strings.Join([]string{"http:/",
                         options.serverAddress,
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

