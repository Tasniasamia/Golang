package cmd;
import (
	"fmt"
	"mains/middleware"
	"net/http"

)
func Start() {
	mux := http.NewServeMux()
	
   managerStruct :=middleware.NewManager();

  managerStruct.Use(middleware.Logger, middleware.Hudai)
   
    globalRoute :=middleware.GlobalRouter(mux);

	InitateRoutes(mux,managerStruct);

	fmt.Println("Server is running on :8080");

	err := http.ListenAndServe(":8080", globalRoute)
	if err != nil {
		fmt.Println("Error starting server:", err)
		panic(err)
	}
}