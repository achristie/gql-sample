package auth

import (
	"log"
	"net/http"
)

func Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("headers: %v", r.Header)
		next.ServeHTTP(w, r)
	}
}
