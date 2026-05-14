package main;

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup;

func uploadFile(user string){
defer wg.Add(-1)
fmt.Println("Start Uploading....");
time.Sleep(3*time.Second);
fmt.Println("Uploaded and return URL");
switch(user){
case "admin":
fmt.Println("Admin");
return;
case "user":
fmt.Println("User");
return;
case "member":
fmt.Println("Member");
return;
}

}

func saveToDB(){

	fmt.Println("Saving url into DB...")
	time.Sleep(3*time.Second);
	fmt.Println("Saved into Database");
	wg.Add(-1)
	
}

func sendEmail(){
fmt.Println("Sending Email....");
time.Sleep(3*time.Second);
fmt.Println("Send Email Successfully");
wg.Add(-1)
}
func main(){
	timeStart :=time.Now();
	wg.Add(1); // counter 1
    go uploadFile("admin");
	wg.Add(1); // counter 2
	go saveToDB();
	wg.Add(1)
	go sendEmail();// counter 3
    wg.Wait();  //wait for counter 0
	fmt.Println("Total time is ",time.Since(timeStart));
}

// used defer because uploadFile(user string) method has no gurantee when it will finished