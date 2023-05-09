package main

import (
    "fmt"
    "testing"

    "net/http/httptest"
    "net/http"
//     "errors"

    "github.com/stretchr/testify/assert"
)

func TestsendReport (t *testing.T) {
    type request struct {
        datatype string
        key any
        value any
    }

    // http server response body
    response := "response"

    tests := []struct {
        name string
        req request
        wanterr error
        wantcode int
    }{
        {
            name: "Test Valid Post request gauge metric",
            req: request{
                datatype: "gauge",
                key: "m01",
                value: 1.34,
            },
            wanterr: nil,
            wantcode: http.StatusOK,
        },
        {
            name: "Test Invalid Post request counter metric",
            req: request{
                datatype: "counter",
                key: "m01",
                value: "ab",
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

            req := server.URL + "/update/" + tc.req.datatype + fmt.Sprintf("/%v/%v", tc.req.key, tc.req.value)
            err := sendReport(req)
            assert.Equal(t, tc.wanterr, err)
        })
    }
}

func TestprocessReport(t *testing.T){
    //to do
}
