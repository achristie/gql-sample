package auth

import (
	"context"
	"net/http"
)

func ApiKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("appkey")
		ctx := context.WithValue(r.Context(), "appkey", header)
		r = r.WithContext(ctx)
		// log.Print(header)
		next.ServeHTTP(w, r)
	})
}
