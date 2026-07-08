package middleware

import (
	"fmt"
	"net/http"
	
)

func Check(mux http.Handler) http.Handler {
	check:=func(w http.ResponseWriter,r *http.Request){
		
		fmt.Println("check Middleware");
		mux.ServeHTTP(w,r);
		
	}

  return http.HandlerFunc(check);
}