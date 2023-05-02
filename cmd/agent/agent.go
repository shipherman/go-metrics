package main

import (
    "errors"
    "fmt"
    "io"
//     "encoding/json"
    "net/http"
    "net/url"
    "runtime"
    "time"
    "math/rand"
//     "strconv"
    s "github.com/shipherman/go-metrics/internal/storage"
)



var stat runtime.MemStats
var pollInterval time.Duration = 2
var reportInterval time.Duration = 10
var m = s.MemStorage{
            GaugeData: make(map[string]s.Gauge),
            CounterData: make(map[string]s.Counter),
    }
var serverAddr = "http://localhost:8080/update/"


func sendPostRequest (dataType string, k, v any) error {
    req := serverAddr + dataType + "/" + fmt.Sprintf("%v/%v",k,v)
    resp, err := http.PostForm(req, url.Values{"Content-type": {"text/plain"}})

    if err != nil {
        return err
    }

//     fmt.Println(resp.Status)
    if resp.StatusCode != http.StatusOK {
        line, err := io.ReadAll(resp.Body)
        if err != nil {
            return err
        }
        return errors.New(fmt.Sprintf(("%s: %s; %s"),
                          "Can't send report to the server",
                          resp.Status,
                          line))
    }
    return nil
}

func processReport (data s.MemStorage) error {
//     fmt.Println(data)
    for k, v := range data.GaugeData {
//         fmt.Println(k,v)
        err := sendPostRequest("gauge", k, v)
        if err != nil {
            return err
        }
    }
    for k, v := range data.CounterData {
//         fmt.Println(k,v)
        err := sendPostRequest("counter", k, v)
        if err != nil {
            return err
        }
    }
    return nil
}

func collectData () {
    runtime.ReadMemStats(&stat)
    m.GaugeData["Alloc"] = s.Gauge(stat.Alloc)
    m.GaugeData["BuckHashSys"] = s.Gauge(stat.BuckHashSys)
    m.GaugeData["Frees"] = s.Gauge(stat.Frees)
    m.GaugeData["GCCPUFraction"] = s.Gauge(stat.GCCPUFraction)
    m.GaugeData["GCSys"] = s.Gauge(stat.GCSys)
    m.GaugeData["HeapAlloc"] = s.Gauge(stat.HeapAlloc)
    m.GaugeData["HeapIdle"] = s.Gauge(stat.HeapIdle)
    m.GaugeData["HeapInuse"] = s.Gauge(stat.HeapInuse)
    m.GaugeData["HeapObjects"] = s.Gauge(stat.HeapObjects)
    m.GaugeData["HeapReleased"] = s.Gauge(stat.HeapReleased)
    m.GaugeData["HeapSys"] = s.Gauge(stat.HeapSys)
    m.GaugeData["LastGC"] = s.Gauge(stat.LastGC)
    m.GaugeData["Lookups"] = s.Gauge(stat.Lookups)
    m.GaugeData["MCacheInuse"] = s.Gauge(stat.MCacheInuse)
    m.GaugeData["MCacheSys"] = s.Gauge(stat.MCacheSys)
    m.GaugeData["MSpanInuse"] = s.Gauge(stat.MSpanInuse)
    m.GaugeData["MSpanSys"] = s.Gauge(stat.MSpanSys)
    m.GaugeData["Mallocs"] = s.Gauge(stat.Mallocs)
    m.GaugeData["NextGC"] = s.Gauge(stat.NextGC)
    m.GaugeData["NumForcedGC"] = s.Gauge(stat.NumForcedGC)
    m.GaugeData["NumGC"] = s.Gauge(stat.NumGC)
    m.GaugeData["OtherSys"] = s.Gauge(stat.OtherSys)
    m.GaugeData["PauseTotalNs"] = s.Gauge(stat.PauseTotalNs)
    m.GaugeData["StackInuse"] = s.Gauge(stat.StackInuse)
    m.GaugeData["StackSys"] = s.Gauge(stat.StackSys)
    m.GaugeData["Sys"] = s.Gauge(stat.Sys)
    m.GaugeData["TotalAlloc"] = s.Gauge(stat.TotalAlloc)
    m.GaugeData["RandomValue"] = s.Gauge(rand.Float32())
    m.CounterData["PollCount"] += 1
}

func main() {
    //collect data through runtime package
    go func() {
        for {
            time.Sleep(time.Second * pollInterval)
            collectData()
        }
    }()

    //send collected data to the server
    for {
        time.Sleep(time.Second * reportInterval)
        err := processReport(m)
        if err != nil {
            panic(err)
        }
    }
}
