package main;

import "fmt";

func count()func()(int){
	count:=0;

	return func()int{
		 count++;
		 return count;
	}

}

func main(){

	counter :=count();
	fmt.Println(counter());//1
	fmt.Println(counter()); //2
	fmt.Println(counter()); //3

}