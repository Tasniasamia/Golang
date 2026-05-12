package main;

import "fmt";



func callBack(a func(b func(c func())),x func(),y func()){
	a(x);
}

func func1(){
 fmt.Println("last functions");
}

func func2(func1 func()){

	func1();

}

func func3(func2 func()){
	func2()
}


func main(){

callBack(func3,func2,func1);
	
	
}