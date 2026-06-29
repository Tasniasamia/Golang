package cmd;
import (
	"fmt"
	"mains/handler"
	"mains/middleware"
	"net/http"
)
func Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handler.GetProducts)
	mux.HandleFunc("POST /product", handler.CreateProduct)
	mux.HandleFunc("GET /product/{id}",handler.GetSingleProduct)

	fmt.Println("Server is running on :8080")
	globalRoute := middleware.GlobalRouter(mux)
	err := http.ListenAndServe(":8080", globalRoute)
	if err != nil {
		fmt.Println("Error starting server:", err)
		panic(err)
	}
}