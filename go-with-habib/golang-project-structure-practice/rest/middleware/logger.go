package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(mux http.Handler) http.Handler {
	log:=func(w http.ResponseWriter,r *http.Request){
		start:=time.Now();
		fmt.Println("Request Method is ",r.Method);
		mux.ServeHTTP(w,r);
		time.Since(start)
	}

  return http.HandlerFunc(log);
}