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

}

func saveToDB(){
	fmt.Println("Saving url into DB...")

	fmt.Println("Saved into Database");
	
}

func sendEmail(){
fmt.Println("Sending Email....");
time.Sleep(3*time.Second);
fmt.Println("Send Email Successfully");

}
func main(){
	timeStart :=time.Now();

    wg.Go(uploadFile); // not working because my golang version is Go 1.22 
    wg.Go(saveToDB); // not working because my golang version is Go 1.22 
    wg.Go(sendEmail); // not working because my golang version is Go 1.22 
    wg.Wait();  //wait for counter 0 
	fmt.Println("Total time is ",time.Since(timeStart));


	
}