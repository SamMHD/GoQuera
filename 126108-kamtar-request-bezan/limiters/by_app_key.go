package limiters

import (
	"database/sql"
	"net/http"

	"golang.org/x/time/rate"
)

func ByAppKey(next http.Handler, refillRate rate.Limit, tokenBucketSize int) http.Handler {
	// TODO: Implement
}
