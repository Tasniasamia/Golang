package middleware

import (
	"log"
	"net/http"
)

func Auth(mux http.Handler) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
     log.Println("Authentication Middleware ",r.Method);
	 mux.ServeHTTP(w,r);
	}
	return http.HandlerFunc(handler);
}
