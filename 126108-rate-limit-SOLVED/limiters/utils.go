package limiters

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

type AccessRecord struct {
	lastAccess   int64
	lastCapacity int64
}

func dumpRequest(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(429)
	fmt.Fprintf(w, `{"error": "too many requests"}`)
}

func updateAccessRecord(token string, new_capacity int64, AccessHistory *map[string]AccessRecord) {
	(*AccessHistory)[token] = AccessRecord{
		lastAccess:   time.Now().Unix(),
		lastCapacity: new_capacity,
	}
}

func ControllAccessByToken(next http.Handler, token string, refillRate rate.Limit, tokenBucketSize int, AccessHistory *map[string]AccessRecord) bool {
	time := time.Now().Unix()
	if access_record, present := (*AccessHistory)[token]; present {
		// TODO: overflow on next line
		newCapacity := access_record.lastCapacity + (time-access_record.lastAccess)*int64(refillRate)
		if newCapacity > int64(tokenBucketSize) {
			newCapacity = int64(tokenBucketSize)
		}
		if newCapacity > 0 {
			updateAccessRecord(token, newCapacity-1, AccessHistory)
			return true
		} else {
			return false
		}
	} else {
		if tokenBucketSize > 0 {
			updateAccessRecord(token, int64(tokenBucketSize)-1, AccessHistory)
			return true
		} else {
			return false
		}
	}

}
