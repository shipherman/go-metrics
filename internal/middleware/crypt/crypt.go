package crypt

import (
	"fmt"
	"bytes"
	"net/http"
	"io/ioutil"
	"encoding/hex"

	"crypto/sha256"
	"crypto/hmac"
)


func CheckReqSign (next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("HashSHA256") == "" {
            fmt.Println("Skip hash calculation")
			next.ServeHTTP(w, r)
            return
        }
		
		var sign []byte

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Calcucalate hash for request body
		h := hmac.New(sha256.New, []byte("key"))
		h.Write(body)
		sha256sum := h.Sum(nil)

		// Read sign from request
		sign, err = hex.DecodeString(r.Header.Get("HashSHA256"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if !hmac.Equal(sign, sha256sum) {
			http.Error(w, "Corrupted sign on request. ", http.StatusBadRequest)
			return
		}

		// Restore request body for further processing 
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		
		next.ServeHTTP(w, r)
	})
}

func SignResponse (next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// To Do
	}
}