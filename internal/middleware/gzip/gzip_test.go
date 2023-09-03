package gzip

import (
	"compress/gzip"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/shipherman/go-metrics/internal/handlers"
	"github.com/shipherman/go-metrics/internal/storage"
)

func TestGzipHandle(t *testing.T) {
	// Init data
	h := handlers.NewHandler()
	h.Store.CounterData = map[string]storage.Counter{"c01": 10}

	gHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			t.Log("Skip compression")
		}

		if strings.Contains(r.Header.Get("Content-Encoding"), "gzip") {
			_, err := gzip.NewReader(r.Body)
			if err != nil {
				t.Error("Could not compress req")
			}
		}
	})
	req := httptest.NewRequest(http.MethodGet, "/value/counter/c01", nil)
	req.Header.Set("Accept-Encoding", "gzip")

	w := httptest.NewRecorder()

	gHandler(w, req)

	gz := GzipHandle(gHandler)
	gz.ServeHTTP(w, req)
}
