package main;

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup;

func uploadFile(){
fmt.Println("Start Uploading....");
time.Sleep(3*time.Second);
fmt.Println("Uploaded and return URL");
wg.Add(-1) //counter 1
}

func saveToDB(){
	fmt.Println("Saving url into DB...")
	time.Sleep(3*time.Second);
	fmt.Println("Saved into Database");
	wg.Add(-1) // counter 0
}

func sendEmail(){
fmt.Println("Sending Email....");
time.Sleep(3*time.Second);
fmt.Println("Send Email Successfully");
wg.Add(-1) // counter 2
}
func main(){
	timeStart :=time.Now();
	wg.Add(1); // counter 1
    go uploadFile();
	wg.Add(1); // counter 2
	go saveToDB();
	wg.Add(1)
	go sendEmail();// counter 3
    wg.Wait();  //wait for counter 0
	fmt.Println("Total time is ",time.Since(timeStart));
}

// output
// Start Uploading....
// Saving url into DB...
// Sending Email....
// Send Email Successfully
// Uploaded and return URL
// Saved into Database
// Total time is  3.002519111s