package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	chi "github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandleUpdate(t *testing.T) {
	type want struct {
		contentType string
		statusCode  int
	}

	type request struct {
		metricType  string
		metricName  string
		metricValue string
	}

	tests := []struct {
		name       string
		request    request
		httpMethod string
		filename   string
		interval   int
		restore    bool
		want       want
	}{
		{
			name: "Test Valid Gauge metric",
			request: request{
				metricType:  "gauge",
				metricName:  "m01",
				metricValue: "1.3",
			},
			httpMethod: http.MethodPost,
			filename:   "/tmp/metrics-db.json",
			interval:   20,
			restore:    true,
			want: want{
				contentType: "text/plain; charset=utf-8",
				statusCode:  http.StatusOK,
			},
		},
		{
			name: "Test Invalid Gauge metric",
			request: request{
				metricType:  "gauge",
				metricName:  "m01",
				metricValue: "1ad3",
			},
			httpMethod: http.MethodPost,
			filename:   "/tmp/metrics-db.json",
			interval:   20,
			restore:    true,
			want: want{
				contentType: "text/plain; charset=utf-8",
				statusCode:  http.StatusBadRequest,
			},
		},
		{
			name: "Test Valid Counter metric",
			request: request{
				metricType:  "counter",
				metricName:  "m02",
				metricValue: "1",
			},
			httpMethod: http.MethodPost,
			filename:   "/tmp/metrics-db.json",
			interval:   20,
			restore:    true,
			want: want{
				contentType: "text/plain; charset=utf-8",
				statusCode:  http.StatusOK,
			},
		},
		{
			name: "Test Invalid Counter metric",
			request: request{
				metricType:  "counter",
				metricName:  "m02",
				metricValue: "1.4",
			},
			httpMethod: http.MethodPost,
			filename:   "/tmp/metrics-db.json",
			interval:   20,
			restore:    true,
			want: want{
				contentType: "text/plain; charset=utf-8",
				statusCode:  http.StatusBadRequest,
			},
		},
		{
			name: "Test Invalid  metric type",
			request: request{
				metricType:  "nosuchmetric",
				metricName:  "m02",
				metricValue: "1.4",
			},
			httpMethod: http.MethodPost,
			filename:   "/tmp/metrics-db.json",
			interval:   20,
			restore:    true,
			want: want{
				contentType: "text/plain; charset=utf-8",
				statusCode:  http.StatusBadRequest,
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			reqString := strings.Join([]string{"/update",
				tc.request.metricType,
				tc.request.metricName,
				tc.request.metricValue}, "/")
			req := httptest.NewRequest(tc.httpMethod, reqString, nil)

			rContext := chi.NewRouteContext()
			rContext.URLParams.Add("type", tc.request.metricType)
			rContext.URLParams.Add("metric", tc.request.metricName)
			rContext.URLParams.Add("value", tc.request.metricValue)

			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rContext))

			h := NewHandler()
			h.HandleUpdate(w, req)

			result := w.Result()

			assert.Equal(t, tc.want.statusCode, result.StatusCode)

			err := result.Body.Close()
			require.NoError(t, err)
		})
	}
}
