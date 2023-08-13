package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"

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
