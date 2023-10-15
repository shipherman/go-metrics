package netchecker

import (
	"errors"
	"fmt"
	"net"
	"net/http"
)

var ErrorEmptyXRealIPHeader = errors.New("empty Header X-Real-IP: ")
var ErrorWrongSubnet = errors.New("wrong subnet: ")

func CheckSubnet(subnet *net.IPNet) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("checking subnet")
			reqIPAddr := r.Header.Get("X-Real-IP")
			if reqIPAddr == "" {
				http.Error(w, ErrorEmptyXRealIPHeader.Error(), http.StatusForbidden)
				return
			}
			// Parse Addr
			ipAddr, _, err := net.ParseCIDR(reqIPAddr)
			if ipAddr == nil {
				http.Error(w, err.Error(), http.StatusForbidden)
				return
			}

			// Check subnet
			if subnet.Contains(ipAddr) {
				next.ServeHTTP(w, r)
			} else {
				fmt.Println(ipAddr)
				http.Error(w, ErrorWrongSubnet.Error(), http.StatusForbidden)
				return
			}
		})
	}
}
