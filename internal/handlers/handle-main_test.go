package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandleMain(t *testing.T) {
	type want struct {
		contentType string
		statusCode  int
	}
	tests := []struct {
		name       string
		request    string
		httpMethod string
		filename   string
		interval   int
		restore    bool
		want       want
	}{
		{
			name:       "Test root page",
			request:    "/",
			httpMethod: http.MethodPost,
			filename:   "/tmp/metrics-db.json",
			interval:   20,
			restore:    true,
			want: want{
				contentType: "text/html; charset=utf-8",
				statusCode:  http.StatusOK,
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.httpMethod, tc.request, nil)
			w := httptest.NewRecorder()

			h := NewHandler()
			h.HandleMain(w, req)

			result := w.Result()
			assert.Equal(t, tc.want.contentType, result.Header.Get("Content-Type"))
			assert.Equal(t, tc.want.statusCode, result.StatusCode)

			err := result.Body.Close()
			require.NoError(t, err)
		})
	}
}
