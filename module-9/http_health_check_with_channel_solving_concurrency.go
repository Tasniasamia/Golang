package main;

import (
	"fmt"
	"net/http"
	"time"
)
type result struct{
	url string;
	status string;
	err error;
}
func checkURL(url string,c chan result){
 res,err :=http.Get(url);
 if(err != nil){
	c <- result{
		url:url,
		status:"Down",
		err:err,

	}
	return;
 }

 c <- result{
url:url,
		status:"Up",
		err:nil,

 }
 defer res.Body.Close();
}

func main(){
    start :=time.Now();
	urlArr :=[] string {
		"https://github.com","https://next.programming-hero.com","https://example.com",
	}

	channel :=make(chan result)
	

 for _,url:=range urlArr{

go checkURL(url,channel)

 }

 for range urlArr{
	result :=channel;
  fmt.Println(<-result);
 }

 fmt.Println("Total time is ",time.Since(start));

}