package main;

import (
	"fmt"
	"time"
);

func uploadFile(c chan string){
fmt.Println("Start Uploading....");
time.Sleep(3*time.Second);
fmt.Println("Uploaded and return URL");
// c <- "https://example.org/image1.png";  //without sending data thread will block and waiting for it at last the thread will deadlock
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
// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [chan receive]:
// main.main()
//         /mnt/c/Users/HP/Desktop/Programming_Hero/Go/module-8/channel/create.channel.deadlock-1.go:20 +0x94
// exit status 2