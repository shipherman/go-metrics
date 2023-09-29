package netchecker

import (
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckSubnet(t *testing.T) {
	type want struct {
		statusCode int
		subnet     string
	}

	tests := []struct {
		name   string
		subnet string
		want   want
	}{
		{
			name:   "Test valid subnet",
			subnet: "192.168.0.1/24",
			want: want{
				statusCode: http.StatusOK,
				subnet:     "192.168.0.0/24",
			},
		},
		{
			name:   "Test Invalid subnet",
			subnet: "192.168.0.1/33",
			want: want{
				statusCode: http.StatusForbidden,
				subnet:     "192.168.0.0/24",
			},
		},
		{
			name:   "Test Wrong subnet",
			subnet: "192.169.0.1/24",
			want: want{
				statusCode: http.StatusForbidden,
				subnet:     "192.168.0.0/24",
			},
		},
		{
			name:   "Test empty real ip header",
			subnet: "",
			want: want{
				statusCode: http.StatusForbidden,
				subnet:     "192.168.0.0/24",
			},
		},
	}

	// Run table tests
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Init data
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			})
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.Header.Set("X-Real-IP", tc.subnet)

			w := httptest.NewRecorder()

			handler(w, req)

			_, subnet, err := net.ParseCIDR(tc.want.subnet)
			assert.NoError(t, err)

			cs := CheckSubnet(subnet)(handler)
			cs.ServeHTTP(w, req)
		})
	}
}
