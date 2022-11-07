package limiters

import (
	"database/sql"
	"fmt"
	"net/http"
	"snapp/db"

	"golang.org/x/time/rate"
)

var appKeyAccessHistory map[string]AccessRecord
var DB *sql.DB

func init() {
	appKeyAccessHistory = make(map[string]AccessRecord)
	DB = db.GetConnection()
}

func IsAppKeyInDB(app_key string) (bool, error) {
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM app_keys WHERE key='%v');", app_key)
	err := DB.QueryRow(query).Scan(&exists)

	return exists, err
}

func ByAppKey(next http.Handler, refillRate rate.Limit, tokenBucketSize int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		appKey := r.Header.Get("X-App-Key")
		fmt.Println("Access By Token", appKey, appKeyAccessHistory[appKey])
		if prs, err := IsAppKeyInDB(appKey); err != nil || !prs {
			next.ServeHTTP(w, r)
		} else {
			if ControllAccessByToken(next, appKey, refillRate, tokenBucketSize, &appKeyAccessHistory) {
				next.ServeHTTP(w, r)
			} else {
				dumpRequest(w)
			}
		}
	})
}
