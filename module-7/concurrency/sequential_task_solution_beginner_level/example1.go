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
	go uploadFile()
    go saveToDB()
	go sendEmail()
	//time.Sleep(4*time.Second); // without it go METHOD_NAME() will not working

	fmt.Println("Total time is ",time.Since(start));
}


/*without time.Sleep() into main function output will be*/

//output
//Total time is  2.223µs




/*With time.Sleep() into main function output wil be*/

// Output
// Sending Email....
// Start Uploading....
// Saving url into DB...
// Send Email Successfully
// Saved into Database
// Uploaded and return URL
// Total time is  6.000478665s