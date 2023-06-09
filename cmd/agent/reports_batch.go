package main

import (
    "io"
    "fmt"
    "strings"
    "net/http"
    "encoding/json"
    "bytes"

    "github.com/shipherman/go-metrics/internal/storage"

)

func sendBatchReport (serverAddress string, metrics []Metrics) error {
    data, err := json.Marshal(metrics)
    if err != nil {
        return err
    }

    data, err = compress(data)
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


func ProcessBatch (serverAddress string, m storage.MemStorage) error {
    var metrics []Metrics

    serverAddress = strings.Join([]string{"http:/",serverAddress,"updates/"}, "/")

    for k, v := range m.CounterData{
        metrics = append(metrics, Metrics{ID:k, MType:counterType, Delta:v})
    }

    for k, v := range m.GaugeData{
        metrics = append(metrics, Metrics{ID:k, MType:gaugeType, Value:v})
    }

    err := sendBatchReport(serverAddress, metrics)
    if err != nil {
            return err
    }
    return nil
}
