package handlers

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleJSONUpdate(t *testing.T) {
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
			request:     "/update/",
			requestBody: `{"id":"c01", "type":"counter", "delta":2}`,
			httpMethod:  http.MethodPost,
			want: want{
				contentType: "application/json",
				statusCode:  http.StatusOK,
			},
		},
		{
			name:        "Test invalid json",
			request:     "/update/",
			requestBody: `{"id":"c01", "type":"incounter", "delta":2}`,
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

			h.HandleJSONUpdate(w, req)

			result := w.Result()

			assert.Equal(t, tc.want.statusCode, result.StatusCode)
		})
	}

}
