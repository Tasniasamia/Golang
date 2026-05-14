package main;

import (
	"fmt"
	"time"
);

func uploadFile(c chan string){
fmt.Println("Start Uploading....");
time.Sleep(3*time.Second);
fmt.Println("Uploaded and return URL");
c <- "https://example.org/image1.png"; 
}


func main(){
start :=time.Now();
channel1 :=make(chan string);
go uploadFile(channel1);
// fmt.Println(<- channel1);
fmt.Println("Total time is ",time.Since(start));
}

// output
// Total time is  1.999µs
// Start Uploading....