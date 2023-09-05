package handlers

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleBatchUpdate(t *testing.T) {
	type want struct {
		contentType string
		statusCode  int
	}

	tests := []struct {
		name        string
		request     string
		requestBody string
		httpMethod  string
		want        want
	}{
		{
			name:        "Test valid json",
			request:     "/updates/",
			requestBody: `[{"id":"c01", "type":"counter", "delta":2}]`,
			httpMethod:  http.MethodPost,
			want: want{
				contentType: "application/json",
				statusCode:  http.StatusOK,
			},
		},
		{
			name:        "Test invalid json",
			request:     "/updates/",
			requestBody: `[{"id":"c01", "type":"incounter", "delta":2}]`,
			httpMethod:  http.MethodPost,
			want: want{
				contentType: "application/json",
				statusCode:  http.StatusBadRequest,
			},
		},
		{
			name:        "Test empty counter metric",
			request:     "/updates/",
			requestBody: `[{"id":"c01", "type":"counter"}]`,
			httpMethod:  http.MethodPost,
			want: want{
				contentType: "application/json",
				statusCode:  http.StatusBadRequest,
			},
		},
		{
			name:        "Test valid gauge metric",
			request:     "/updates/",
			requestBody: `[{"id":"c01", "type":"gauge", "value":3.13}]`,
			httpMethod:  http.MethodPost,
			want: want{
				contentType: "application/json",
				statusCode:  http.StatusOK,
			},
		},
		{
			name:        "Test empty gauge metric",
			request:     "/updates/",
			requestBody: `[{"id":"c01", "type":"gauge"}]`,
			httpMethod:  http.MethodPost,
			want: want{
				contentType: "application/json",
				statusCode:  http.StatusBadRequest,
			},
		},
		{
			name:        "Test corrupted json",
			request:     "/updates/",
			requestBody: `[{"id":"c01", "type":"gauge"},]`,
			httpMethod:  http.MethodPost,
			want: want{
				contentType: "application/json",
				statusCode:  http.StatusBadRequest,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			h := NewHandler()

			w := httptest.NewRecorder()
			req := httptest.NewRequest(tc.httpMethod, tc.request, nil)
			req.Body = io.NopCloser(bytes.NewBuffer([]byte(tc.requestBody)))

			h.HandleBatchUpdate(w, req)

			result := w.Result()
			assert.Equal(t, tc.want.statusCode, result.StatusCode)

			result.Body.Close()
		})
	}
}
