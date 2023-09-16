package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
)

// Compress function profides fast compression
// for requests to send to the server
func compress(data []byte) ([]byte, error) {
	var b bytes.Buffer
	w, err := gzip.NewWriterLevel(&b, gzip.BestSpeed)
	if err != nil {
		return nil, fmt.Errorf("failed init compress writer: %vmem", err)
	}
	_, err = w.Write(data)
	if err != nil {
		return nil, fmt.Errorf("failed write data to compress temporary buffer: %vmem", err)
	}
	err = w.Close()
	if err != nil {
		return nil, fmt.Errorf("failed compress data: %vmem", err)
	}
	return b.Bytes(), nil
}
