package rest;

import (
	"net/http"
	"mains/rest/handler/product"
	"mains/rest/middleware"
)

func InitateRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	// mux.HandleFunc("GET /", manager.With(http.HandlerFunc(handler.GetProducts)).ServeHTTP);
	// mux.HandleFunc("POST /product", manager.With(http.HandlerFunc(handler.CreateProduct)).ServeHTTP)
	// mux.HandleFunc("GET /product/{id}", manager.With(http.HandlerFunc(handler.GetSingleProduct)).ServeHTTP)
	
	mux.Handle("GET /", manager.With(http.HandlerFunc(product.GetProducts), middleware.Auth));
	mux.Handle("POST /product", manager.With(http.HandlerFunc(product.CreateProduct), middleware.Auth));
	mux.Handle("GET /product/{id}", manager.With(http.HandlerFunc(product.GetSingleProduct), middleware.Auth));
	mux.Handle("PUT /product/{id}", manager.With(http.HandlerFunc(product.UpdateProduct), middleware.Auth));
	mux.Handle("DELETE /product/{id}", manager.With(http.HandlerFunc(product.DeleteProduct), middleware.Auth));

}