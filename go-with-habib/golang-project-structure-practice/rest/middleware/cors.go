package middleware

import (
	"fmt"
	"net/http"
)

func Cors(mux http.Handler) http.Handler {
	log:=func(w http.ResponseWriter,r *http.Request){
	w.Header().Add("Access-Control-Allow-Origin", "*");
    w.Header().Add("Access-Control-Allow-Headers", "Content-Type,Authorization");
	w.Header().Add("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS");
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	fmt.Println("Request Method is ",r.Method);
	mux.ServeHTTP(w,r);
	}

  return http.HandlerFunc(log);
}