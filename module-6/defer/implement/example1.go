package main;

import "fmt";



func deferFunc()(int){
	a :=10;
	defer func(){
    fmt.Println("The value of a is ",a);
	}();

	fmt.Println("After defer code");
	a+=10;
	return a;
}


func realDefer(a int)(int){

	return a*100;

}
func deferFunc2()(int){
	a:=0;
	fmt.Println("Start here");
   defer fmt.Println(realDefer(a));

	fmt.Println("After defer code");
	a+=10;
	return a;
}



func main(){
fmt.Println(deferFunc());
fmt.Printf("\n \n");
fmt.Println(deferFunc2());

}