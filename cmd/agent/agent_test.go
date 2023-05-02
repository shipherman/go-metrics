package main

import (
    "testing"
)

func TestSendPostRequest (t *testing.T) {
    type request struct {
        datatype string
        key any
        value any
    }
    tests := []struct {
        name string
        req request
        wanterr error
    }{
        {
            name: "Test Valid Post request gauge metric",
            req: req{
                datatype: "gauge",
                key: "m01",
                value: 1.34,
            },
        },
    }
}
