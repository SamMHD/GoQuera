package limiters

import (
	"fmt"
	"net/http"

	"golang.org/x/time/rate"
)

var ipv4AccessHistory map[string]AccessRecord

func init() {
	ipv4AccessHistory = make(map[string]AccessRecord)
}

func ByIp(next http.Handler, refillRate rate.Limit, tokenBucketSize int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ip := r.RemoteAddr
		fmt.Println("Access By IP", ip, ipv4AccessHistory[ip])

		if ControllAccessByToken(next, ip, refillRate, tokenBucketSize, &ipv4AccessHistory) {
			next.ServeHTTP(w, r)
		} else {
			dumpRequest(w)
		}

	})
}
