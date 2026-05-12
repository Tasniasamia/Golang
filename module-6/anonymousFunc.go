package main;

import "fmt";

func callback(a func(x int,y int)int,m int,n int)int{
	return a(m,n);
}

func main(){
	fmt.Println(callback(func(a int,b int)int{return a+b;},3,4));
	
}