

package main;

import "fmt";

func getData(user string,roll int)(userName string,rollNumber int){
userName =user;
rollNumber =roll;
return;
	
}

func getData2(user string,roll int){
	fmt.Println("user is ",user);
	fmt.Println("roll is ",roll);
}

func deferFunc(name string,roll int)(int){
	
	fmt.Println("Start Function");

     /* That codes of 20 number line to 24 number  will give error . because defer will never return */
	// defer userName,rollNumber :=getData(name,roll);
    // fmt.Println(userName);
	// fmt.Println(rollNumber);
    //  rollNumber =43;
	//  return rollNumber;
    
	userName,rollNumber :=getData(name,roll)
	


	defer getData2(userName,rollNumber);

	rollNumber=22;

	return rollNumber;

}

func main(){

fmt.Println(deferFunc("Tasnia",1));

}


//output
// Start Function
// user is  Tasnia
// roll is  1
// 22

//It's a wrong output format right? but no 
// rollNumber already returned before defer . then defer worked then into function fmt.Println will worked. 