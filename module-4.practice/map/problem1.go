package main;

import "fmt";

func main(){
	student:=map[string]int {
		"samia":80,
		"lamia":72,
		"disha":23,
	}
    fmt.Println(student);

	student["samia"]=90;
	fmt.Println(student);

	delete(student,"lamia");
	fmt.Println(student);

	for key,value := range student{
		fmt.Println(key,"=",value);
	}


}