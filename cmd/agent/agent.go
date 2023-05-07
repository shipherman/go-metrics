package main

import (
    "fmt"
    "io"
    "log"
    "flag"
    "runtime"
    "time"
    "math/rand"
    "strconv"

    "net/http"
    "net/url"

    s "github.com/shipherman/go-metrics/internal/storage"
//     e "github.com/caarlos0/env/v6"

    "os"
)

//cli options
type Options struct {
    serverAddress string `env:"ADDRESS"`
    pollInterval int `env:"POLL_INTERVAL"`
    reportInterval int `env:"REPORT_INTERVAL"`
}

var options Options

//init MemStorage
var m = s.MemStorage{Data: make(map[string]interface{})}

//MemStats instance
var stat runtime.MemStats

//server parameters
var contentType = url.Values{"Content-type": {"text/plain"}}




var logger *log.Logger


func parseOptions () {
// 	fmt.Println(options)
    flag.IntVar(&options.pollInterval, "p", 2,
                     "Frequensy in seconds for collecting metrics")
    flag.IntVar(&options.reportInterval, "r", 10,
                     "Frequensy in seconds for sending report to the server")
    flag.StringVar(&options.serverAddress, "a", "localhost:8080",
                "Address of the server to send metrics")
    flag.Parse()
// 	fmt.Println(options)

    // env/v6 не работает )Ж
//     if err := e.Parse(&options); err != nil {
//         fmt.Println(err)
//     }

    if l := os.Getenv("ADDRESS"); l != "" {
        options.serverAddress = l
    }
    if l := os.Getenv("POLL_INTERVAL"); l != "" {
        i, err := strconv.Atoi(l)
        if err != nil {
            panic(err)
        }
        options.pollInterval = i
    }
    if l := os.Getenv("REPORT_INTERVAL"); l != "" {
        i, err := strconv.Atoi(l)
        if err != nil {
            panic(err)
        }
        options.reportInterval = i
    }

// 	fmt.Println(options, os.Getenv("POLL_INTERVAL"), os.Getenv("REPORT_INTERVAL") )
}

func SendPostRequest (req string) error {
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
        req := "http://" + options.serverAddress + "/update/" + mtype + fmt.Sprintf("/%v/%v",k,v)
        err := SendPostRequest(req)
        if err != nil {
            return err
        }
    }
    return nil
}

func main() {
    //parse cli options
    parseOptions()

    // initiate conters
    m.Data["PollCount"] = s.Counter(0)

    go func() {
        for {
            //collect data from MemStats
            runtime.ReadMemStats(&stat)
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

            //collect timeout
            time.Sleep(time.Duration(options.pollInterval) * time.Second)

            }
        }()

    //send collected data to the server
    for {
        err := ProcessReport(m)
        if err != nil {
            log.Println(err)
        }
        //report timeout
        time.Sleep(time.Duration(options.reportInterval) * time.Second)

    }
}
