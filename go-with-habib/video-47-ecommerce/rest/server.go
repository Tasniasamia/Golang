package rest;

import (
	"fmt"
	"mains/config"
	"mains/rest/middleware"
	"net/http"

)
func Start() {
   mux := http.NewServeMux()
	
   managerStruct :=middleware.NewManager();

   managerStruct.Use(middleware.Logger, middleware.Hudai,middleware.Preflight,middleware.Cors)
   
   globalRoute :=managerStruct.WrappedMux(mux);



	InitateRoutes(mux,managerStruct);

	fmt.Println("Server is running on :8080");
    address :=":"+config.GetConfig().HTTP_PORT
	err := http.ListenAndServe(address, globalRoute)
	if err != nil {
		fmt.Println("Error starting server:", err)
		panic(err)
	}
}