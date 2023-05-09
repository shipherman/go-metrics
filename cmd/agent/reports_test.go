package main

import (
    "fmt"
    "testing"
    "strings"
    "net/http/httptest"
    "net/http"
//     "errors"

    "github.com/stretchr/testify/assert"
    s "github.com/shipherman/go-metrics/internal/storage"
)

func TestProcessReport (t *testing.T) {
    // http server response body
    response := "response"

    tests := []struct {
        name string
        store s.MemStorage
        wanterr error
        wantcode int
    }{
        {
            name: "Test Valid Post request gauge metric",
            store: s.MemStorage{
                Data: map[string]interface{}{
                    "valid": s.Gauge(2.32),
                },
            },
            wanterr: nil,
            wantcode: http.StatusOK,
        },
        {
            name: "Test Invalid Post request counter metric",
             store: s.MemStorage{
                Data: map[string]interface{}{
                    "valid": s.Counter(2),
                },
            },
            // adding new line into format string as http server do
            wanterr: fmt.Errorf("%s: %s; %s\n",
                                "Can't send report to the server",
                                "400 Bad Request",
                                response),
            wantcode: http.StatusBadRequest,
        },
    }
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T){
            server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
                http.Error(rw, response, tc.wantcode)
            }))
            defer server.Close()

            options.serverAddress = strings.Replace(server.URL, "http://", "", 1)
            err := ProcessReport(&tc.store)
            assert.Equal(t, tc.wanterr, err)
        })
    }
}

