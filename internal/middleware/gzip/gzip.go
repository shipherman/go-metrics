package gzip

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

type gzipWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

// Zipper middleware
func GzipHandle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		gz, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		defer gz.Close()

		if strings.Contains(r.Header.Get("Content-Encoding"), "gzip") {
			r.Body, err = gzip.NewReader(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		w.Header().Set("Content-Encoding", "gzip")
		next.ServeHTTP(gzipWriter{ResponseWriter: w, Writer: gz}, r)
	})
}
