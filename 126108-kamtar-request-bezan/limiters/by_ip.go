package limiters

import (
	"net/http"

	"golang.org/x/time/rate"
)

func ByIp(next http.Handler, refillRate rate.Limit, tokenBucketSize int) http.Handler {
	// TODO: Implement
}
