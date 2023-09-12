package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shipherman/go-metrics/internal/storage"

	chi "github.com/go-chi/chi/v5"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// BenchMark Update request handling
func BenchmarkHandleUpdate(b *testing.B) {
	w := httptest.NewRecorder()

	req := httptest.NewRequest("POST", "/update/gauge/m01/1.3", nil)

	rContext := chi.NewRouteContext()
	rContext.URLParams.Add("type", "gauge")
	rContext.URLParams.Add("metric", "m01")
	rContext.URLParams.Add("value", "1.3")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rContext))

	h := NewHandler()
	h.HandleUpdate(w, req)

	result := w.Result()

	assert.Equal(b, http.StatusOK, result.StatusCode)

	err := result.Body.Close()
	require.NoError(b, err)

}

// BenchMark Value request handling
func BenchmarkHandleValue(b *testing.B) {
	w := httptest.NewRecorder()

	req := httptest.NewRequest("GET", "/update/gauge/m01", nil)

	rContext := chi.NewRouteContext()
	rContext.URLParams.Add("metric", "m01")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rContext))

	h := NewHandler()
	h.Store.UpdateGauge("m01", 1.3)
	h.HandleValue(w, req)

	result := w.Result()

	assert.Equal(b, http.StatusOK, result.StatusCode)

	err := result.Body.Close()
	require.NoError(b, err)

}

// BenchMark Batch handler
func BenchmarkHandleBatchUpdate(b *testing.B) {
	type Metrics struct {
		ID    string          `json:"id"`    // имя метрики
		MType string          `json:"type"`  // параметр, принимающий значение gauge или counter
		Delta storage.Counter `json:"delta"` // значение метрики в случае передачи counter
		Value storage.Gauge   `json:"value"` // значение метрики в случае передачи gauge
	}

	metrics := []Metrics{
		{
			ID:    "m01",
			MType: "gauge",
			Value: 3.14,
		},
		{
			ID:    "m02",
			MType: "counter",
			Delta: 20,
		},
	}

	h := NewHandler()

	w := httptest.NewRecorder()

	req := httptest.NewRequest("POST", "/updates/", nil)
	data, _ := json.Marshal(metrics)
	req.Body = io.NopCloser(bytes.NewBuffer(data))
	h.HandleBatchUpdate(w, req)

	result := w.Result()

	assert.Equal(b, http.StatusOK, result.StatusCode)

	err := result.Body.Close()
	require.NoError(b, err)

}
