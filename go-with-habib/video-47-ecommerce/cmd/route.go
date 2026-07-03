package cmd;

import (
	"net/http"
	"mains/handler"
	"mains/middleware"
)

func InitateRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	// mux.HandleFunc("GET /", manager.With(http.HandlerFunc(handler.GetProducts)).ServeHTTP);
	// mux.HandleFunc("POST /product", manager.With(http.HandlerFunc(handler.CreateProduct)).ServeHTTP)
	// mux.HandleFunc("GET /product/{id}", manager.With(http.HandlerFunc(handler.GetSingleProduct)).ServeHTTP)
	
	mux.Handle("GET /", manager.With(http.HandlerFunc(handler.GetProducts), middleware.Auth));
	mux.Handle("POST /product", manager.With(http.HandlerFunc(handler.CreateProduct), middleware.Auth));
	mux.Handle("GET /product/{id}", manager.With(http.HandlerFunc(handler.GetSingleProduct), middleware.Auth));


	
}