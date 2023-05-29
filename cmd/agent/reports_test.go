package agent

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
    responseBody := "response"

    tests := []struct {
        name string
        store s.MemStorage
        wanterr error
        wantcode int
    }{
        {
            name: "Test Valid Post request gauge metric",
            store: s.MemStorage{
                GaugeData: map[string]s.Gauge{
                    "valid": s.Gauge(2.32),
                },
            },
            wanterr: nil,
            wantcode: http.StatusOK,
        },
        {
            name: "Test Empty metric",
            store: s.MemStorage{CounterData: map[string]s.Counter{}},
            // adding new line into format string as http server do
            wanterr: nil,
            wantcode: http.StatusBadRequest,
        },
        {
            name: "Test Invalid Post request counter metric",
             store: s.MemStorage{
                CounterData: map[string]s.Counter{
                    "valid": s.Counter(2),
                },
            },
            // adding new line into format string as http server do
            wanterr: fmt.Errorf("%s: %s; %s\n",
                                "Can't send report to the server",
                                "400 Bad Request",
                                responseBody),
            wantcode: http.StatusBadRequest,
        },
    }
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T){
            server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
                http.Error(rw, responseBody, tc.wantcode)
            }))
            defer server.Close()

//             .serverAddress =
            err := ProcessReport(strings.Replace(server.URL, "http://", "", 1), tc.store)
            assert.Equal(t, tc.wanterr, err)
        })
    }
}

