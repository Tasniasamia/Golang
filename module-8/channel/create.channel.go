package main;

import (
	"fmt"
	"time"
);

func uploadFile(c chan string){
fmt.Println("Start Uploading....");
time.Sleep(3*time.Second);
fmt.Println("Uploaded and return URL");
c <- "https://example.org/image1.png";  //send channel value
}


func main(){
	start :=time.Now();
 channel1 :=make(chan string);
go uploadFile(channel1);
fmt.Println(<- channel1); //recive channel value
fmt.Println("Total time is ",time.Since(start));

}

// output
// Start Uploading....
// Uploaded and return URL
// https://example.org/image1.png
// Total time is  3.001727031s