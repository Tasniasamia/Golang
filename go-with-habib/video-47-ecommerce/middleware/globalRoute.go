package middleware;
import "net/http"

func GlobalRouter(mux *http.ServeMux) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*");
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type");
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS");
	w.Header().Set("Content-Type","application/json");
		if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	 mux.ServeHTTP(w,r);
	}
	return http.HandlerFunc(handler);
}
