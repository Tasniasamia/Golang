package main;

import (
"fmt"
"time"
) 

func uploadFile(){
fmt.Println("Start Uploading....");
time.Sleep(3*time.Second);
fmt.Println("Uploaded and return URL");
}

func saveToDB(){
	fmt.Println("Saving url into DB...")
	time.Sleep(3*time.Second);
	fmt.Println("Saved into Database");
}

func sendEmail(){
fmt.Println("Sending Email....");
time.Sleep(3*time.Second);
fmt.Println("Send Email Successfully");
}

func main(){
	start :=time.Now();
	uploadFile()
    saveToDB()
	sendEmail()
	fmt.Println("Total time is ",time.Since(start));
}