package logger

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shipherman/go-metrics/internal/handlers"
	"github.com/shipherman/go-metrics/internal/storage"
)

func TestLogHandler(t *testing.T) {
	h := handlers.NewHandler()
	h.Store.CounterData = map[string]storage.Counter{"c01": 10}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	req := httptest.NewRequest(http.MethodGet, "/value/counter/c01", nil)

	w := httptest.NewRecorder()

	handler(w, req)

	gz := LogHandler(handler)
	gz.ServeHTTP(w, req)
}
