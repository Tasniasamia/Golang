package main;

import (
	"fmt"
	"net/http"
	"time"
)


func checkURL(url string){
 res,err :=http.Get(url);
 if(err != nil){
    fmt.Println("url is ",url);
	fmt.Println("status is Down");
	fmt.Println("error is",err);
	return;
 }

    fmt.Println("url is ",url);
	fmt.Println("status is UP");
	fmt.Println("error is",err);
 defer res.Body.Close();
}

func main(){

	start :=time.Now();
    
	urlArr :=[] string {
		"https://github.com","https://next.programming-hero.com","https://example.com",
	}

	

for _,url:=range urlArr{
 checkURL(url)
}


 fmt.Println("total time is ",time.Since(start));

}