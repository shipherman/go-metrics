package handlers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/shipherman/go-metrics/internal/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandlePing(t *testing.T) {
	cfgDSN := os.Getenv("DATABASE_DSN")

	type want struct {
		statusCode int
		err        error
	}
	tests := []struct {
		name       string
		request    string
		httpMethod string
		dsn        string
		want       want
	}{
		{
			name:       "Test ping page - OK",
			request:    "/ping",
			httpMethod: http.MethodGet,
			dsn:        cfgDSN,
			want: want{
				statusCode: http.StatusOK,
				err:        nil,
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.httpMethod, tc.request, nil)
			w := httptest.NewRecorder()

			h := NewHandler()

			database, _ := db.Connect(tc.dsn)
			h.DBconn = database.Conn

			h.HandlePing(w, req)

			result := w.Result()
			assert.Equal(t, tc.want.statusCode, result.StatusCode)

			err := result.Body.Close()
			require.NoError(t, err)
		})
	}
}
