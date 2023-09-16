package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/shipherman/go-metrics/internal/storage"
)

func sendBatchReport(cfg Options, metrics []Metrics) error {
	// var sha256sum string

	data, err := json.Marshal(metrics)
	if err != nil {
		return err
	}
	// fmt.Println("raw data size: ", len(data))

	// Init request
	request, err := http.NewRequest("POST", cfg.ServerAddress, bytes.NewBuffer([]byte{}))
	if err != nil {
		return err
	}

	data, err = compress(data)
	if err != nil {
		return err
	}
	// fmt.Println("compressed data size: ", len(data))

	// Encrypt data and set Header
	// if cfg.Encrypt {
	// 	h := hmac.New(sha256.New, cfg.KeyByte)
	// 	h.Write(data)
	// 	sha256sum = hex.EncodeToString(h.Sum(nil))
	// 	request.Header.Set("HashSHA256", sha256sum)
	// }
	if cfg.Encrypt {
		data, err = Encrypt(cfg.CryptoKey, data)
		if err != nil {
			return err
		}
	}

	// Redefine request content
	request.Body = io.NopCloser(bytes.NewBuffer(data))

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

// Marshal metrics slice into JSON then send it to the server via sendBathReport() func.
func ProcessBatch(ctx context.Context, cfg Options,
	metricsCh chan storage.MemStorage) error {
	var metrics []Metrics

	// Receive MemStorage with actual metrics
	m := <-metricsCh

	// Prepare structure to send to the server
	for k, v := range m.CounterData {
		metrics = append(metrics, Metrics{ID: k, MType: counterType, Delta: v})
	}
	for k, v := range m.GaugeData {
		metrics = append(metrics, Metrics{ID: k, MType: gaugeType, Value: v})
	}

	// Send report
	err := sendBatchReport(cfg, metrics)
	if err != nil {
		return err
	}

	return nil
}
