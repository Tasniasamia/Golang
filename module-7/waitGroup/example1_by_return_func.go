package main;

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup;
var role string;

func uploadFile(user string){
defer wg.Add(-1)
fmt.Println("Start Uploading....");
time.Sleep(3*time.Second);
fmt.Println("Uploaded and return URL");
switch(user){
case "admin":
role="admin"
return;
case "user":
role="user";
return;
case "member":
role="member";
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
   //fmt.Println("user role is ",role); // not give us proper output because uploadfile() goRoutine cannot be end that time.
	wg.Add(1); // counter 2
	go saveToDB();
	wg.Add(1)
	go sendEmail();// counter 3
    wg.Wait();  //wait for counter 0
	fmt.Println("user role is ",role);
	fmt.Println("Total time is ",time.Since(timeStart));
}

// goRoutine never return function we used here global variable for getting the output of function.but it's not a good feature
