package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	handler:= http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start:=time.Now();
		log.Println("Request:", r.Method, r.URL.Path);
		next.ServeHTTP(w, r);
		end :=time.Since(start);
		log.Println("Response time:", end);
	})
	return handler
}