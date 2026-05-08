package main;
import "fmt";

func modifyWithPointer(arr *[5] int){

// (*arr)[0]=122; // এটাও same কাজ করে — Go auto dereference করে
// (*arr)[4]=555; // এটাও same কাজ করে — Go auto dereference করে
arr[0]=122;
arr[4]=555;
fmt.Println(*arr);
}


func main(){

	arr:= [5]int{1,2,3,4,5};

	modifyWithPointer(&arr);
	fmt.Println("outside of function",arr);
	fmt.Println("outside of function2",arr);



	a:=12;
	b:=&a;
	a=133;
	fmt.Println("a=",a);
	fmt.Println("b=",b);


}