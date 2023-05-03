package main

import (
//     "fmt"
    "testing"
    "net/http"
    "net/http/httptest"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"


)


func TestHandleMain (t *testing.T){
    type want struct {
        contentType string
        statusCode int
    }
    tests := []struct{
        name string
        request string
        httpMethod string
        want want
    }{
        {
            name: "Test root page",
            request: "/",
            httpMethod: http.MethodPost,
            want: want{
                contentType: "text/plain; charset=utf-8",
                statusCode: http.StatusBadRequest,
            },
        },
    }
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T){
            request := httptest.NewRequest(tc.httpMethod, tc.request, nil)
            w := httptest.NewRecorder()
            HandleMain(w, request)

            result := w.Result()
            assert.Equal(t, tc.want.contentType, result.Header.Get("Content-Type"))
            assert.Equal(t, tc.want.statusCode, result.StatusCode)

            err := result.Body.Close()
            require.NoError(t, err)
        })
    }
}

func TestHandleUpdate (t *testing.T) {
    type want struct {
        contentType string
        statusCode int
    }
    tests := []struct{
        name string
        request string
        httpMethod string
        want want
    }{
        {
            name: "Test Valid Gauge metric",
            request: "/update/gauge/m01/1.35",
            httpMethod: http.MethodPost,
            want: want{
                contentType: "text/plain; charset=utf-8",
                statusCode: http.StatusOK,
            },
        },
        {
            name: "Test Invalid Gauge metric",
            request: "/update/gauge/m02/1e",
            httpMethod: http.MethodPost,
            want: want{
                contentType: "text/plain; charset=utf-8",
                statusCode: http.StatusBadRequest,
            },
        },
        {
            name: "Test Valid Counter metric",
            request: "/update/counter/n01/5",
            httpMethod: http.MethodPost,
            want: want{
                contentType: "text/plain; charset=utf-8",
                statusCode: http.StatusOK,
            },
        },
        {
            name: "Test Invalid Counter metric",
            request: "/update/counter/n01/5.2",
            httpMethod: http.MethodPost,
            want: want{
                contentType: "text/plain; charset=utf-8",
                statusCode: http.StatusBadRequest,
            },
        },
        {
            name: "Test Invalid update request without metric name",
            request: "/update/gauge/12.3",
            httpMethod: http.MethodPost,
            want: want{
                contentType: "text/plain; charset=utf-8",
                statusCode: http.StatusNotFound,
            },
        },
    }
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T){
            request := httptest.NewRequest(tc.httpMethod, tc.request, nil)
            w := httptest.NewRecorder()
            HandleUpdate(w, request)

            result := w.Result()
            assert.Equal(t, tc.want.statusCode, result.StatusCode)

            err := result.Body.Close()
            require.NoError(t, err)
        })
    }
}
