package main

import (
    "fmt"
    "io"
    "runtime"
    "math/rand"
//     "encoding/json"
    "net/http"
    "net/url"
    "time"

    s "github.com/shipherman/go-metrics/internal/storage"
)


//sleep param
var pollInterval time.Duration = 2
var reportInterval time.Duration = 10

//init MemStorage
var m = s.MemStorage{Data: make(map[string]interface{})}

//server parameters
var serverAddr = "http://localhost:8080/update/"
var contentType = url.Values{"Content-type": {"text/plain"}}

func SendPostRequest (req string) error {
    //build request string

    resp, err := http.PostForm(req, contentType)
    if err != nil {
        return err
    }
    //     fmt.Println(resp.Status)
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

func ProcessReport (data s.MemStorage) error {
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
        req := serverAddr + mtype + fmt.Sprintf("/%v/%v",k,v)
        err := SendPostRequest(req)
        if err != nil {
            return err
        }
    }
    return nil
}

func main() {
    var stat runtime.MemStats
    runtime.ReadMemStats(&stat)

    // initiate conters
    m.Data["PollCount"] = s.Counter(0)

    go func() {
        for {
            time.Sleep(time.Second * pollInterval)

            //collect data from MemStats
            m.Data["Alloc"] = s.Gauge(stat.Alloc)
            m.Data["BuckHashSys"] = s.Gauge(stat.BuckHashSys)
            m.Data["Frees"] = s.Gauge(stat.Frees)
            m.Data["GCCPUFraction"] = s.Gauge(stat.GCCPUFraction)
            m.Data["GCSys"] = s.Gauge(stat.GCSys)
            m.Data["HeapAlloc"] = s.Gauge(stat.HeapAlloc)
            m.Data["HeapIdle"] = s.Gauge(stat.HeapIdle)
            m.Data["HeapInuse"] = s.Gauge(stat.HeapInuse)
            m.Data["HeapObjects"] = s.Gauge(stat.HeapObjects)
            m.Data["HeapReleased"] = s.Gauge(stat.HeapReleased)
            m.Data["HeapSys"] = s.Gauge(stat.HeapSys)
            m.Data["LastGC"] = s.Gauge(stat.LastGC)
            m.Data["Lookups"] = s.Gauge(stat.Lookups)
            m.Data["MCacheInuse"] = s.Gauge(stat.MCacheInuse)
            m.Data["MCacheSys"] = s.Gauge(stat.MCacheSys)
            m.Data["MSpanInuse"] = s.Gauge(stat.MSpanInuse)
            m.Data["MSpanSys"] = s.Gauge(stat.MSpanSys)
            m.Data["Mallocs"] = s.Gauge(stat.Mallocs)
            m.Data["NextGC"] = s.Gauge(stat.NextGC)
            m.Data["NumForcedGC"] = s.Gauge(stat.NumForcedGC)
            m.Data["NumGC"] = s.Gauge(stat.NumGC)
            m.Data["OtherSys"] = s.Gauge(stat.OtherSys)
            m.Data["PauseTotalNs"] = s.Gauge(stat.PauseTotalNs)
            m.Data["StackInuse"] = s.Gauge(stat.StackInuse)
            m.Data["StackSys"] = s.Gauge(stat.StackSys)
            m.Data["Sys"] = s.Gauge(stat.Sys)
            m.Data["TotalAlloc"] = s.Gauge(stat.TotalAlloc)
            m.Data["RandomValue"] = s.Gauge(rand.Float32())
            m.Data["PollCount"] = m.Data["PollCount"].(s.Counter) + 1
            }
        }()

    //send collected data to the server
    for {
        time.Sleep(time.Second * reportInterval)
        err := ProcessReport(m)
        if err != nil {
            panic(err)
        }
    }
}
