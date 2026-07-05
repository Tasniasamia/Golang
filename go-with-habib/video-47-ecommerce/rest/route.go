package rest;

import (
	"net/http"
	"mains/rest/handler/product"
	"mains/rest/handler/user"
	"mains/rest/middleware"

)

func InitateRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	// mux.HandleFunc("GET /", manager.With(http.HandlerFunc(handler.GetProducts)).ServeHTTP);
	// mux.HandleFunc("POST /product", manager.With(http.HandlerFunc(handler.CreateProduct)).ServeHTTP)
	// mux.HandleFunc("GET /product/{id}", manager.With(http.HandlerFunc(handler.GetSingleProduct)).ServeHTTP)
	
	mux.Handle("GET /", manager.With(http.HandlerFunc(product.GetProducts)));
	mux.Handle("POST /product", manager.With(http.HandlerFunc(product.CreateProduct), middleware.AuthMiddleware));
	mux.Handle("GET /product/{id}", manager.With(http.HandlerFunc(product.GetSingleProduct)));
	mux.Handle("PUT /product/{id}", manager.With(http.HandlerFunc(product.UpdateProduct), middleware.AuthMiddleware));
	mux.Handle("DELETE /product/{id}", manager.With(http.HandlerFunc(product.DeleteProduct), middleware.AuthMiddleware));
	mux.Handle("POST /resister", manager.With(http.HandlerFunc(user.CreateUser)));
	mux.Handle("POST /login", manager.With(http.HandlerFunc(user.Login)));

}