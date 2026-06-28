package main;

import (
	"fmt"
   "net/http"
)
func rootHandler(res http.ResponseWriter,req * http.Request){
 	fmt.Fprintln(res, "welcom to go sever")

}
func main(){
mux :=http.NewServeMux();
mux.HandleFunc("/",rootHandler);
err :=http.ListenAndServe(":3030",mux);
if (err != nil){
	panic(err);
}
}