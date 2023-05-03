package main

import (
    "errors"
    "fmt"
    "io"
//     "encoding/json"
    "net/http"
    "net/url"
//     "runtime"
    "time"
//     "math/rand"
//     "strconv"
    s "github.com/shipherman/go-metrics/internal/storage"
)


var pollInterval time.Duration = 2
var reportInterval time.Duration = 10
var m = s.MemStorage{
            GaugeData: make(map[string]s.Gauge),
            CounterData: make(map[string]s.Counter),
    }
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
        return errors.New(fmt.Sprintf(("%s: %s; %s"),
                          "Can't send report to the server",
                          resp.Status,
                          line))
    }
    return nil
}

func ProcessReport (data s.MemStorage) error {
//     fmt.Println(data)
    for k, v := range data.GaugeData {
//         fmt.Println(k,v)
        req := serverAddr + "gauge" + fmt.Sprintf("/%v/%v",k,v)
        err := SendPostRequest(req)
        if err != nil {
            return err
        }
    }
    for k, v := range data.CounterData {
//         fmt.Println(k,v)
        req := serverAddr + "counter" + fmt.Sprintf("/%v/%v",k,v)
        err := SendPostRequest(req)
        if err != nil {
            return err
        }
    }
    return nil
}

func main() {
    //collect data through runtime package
    go func() {
        for {
            time.Sleep(time.Second * pollInterval)
            s.UpdateMemStorage(&m)
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
