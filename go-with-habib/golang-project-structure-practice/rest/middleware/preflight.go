package middleware

import (
	"fmt"
	"net/http"
	
)

func Preflight(mux http.Handler) http.Handler {
	log:=func(w http.ResponseWriter,r *http.Request){

		if(r.Method == http.MethodOptions){
			w.WriteHeader(http.StatusAccepted);
			return;
		}
		
		fmt.Println("Request Method is ",r.Method);
		mux.ServeHTTP(w,r);
		
	}

  return http.HandlerFunc(log);
}