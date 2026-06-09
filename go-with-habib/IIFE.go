package main;
import (
	"fmt"
)

func main(){
	func(a,b int){
		fmt.Println(a+b)
	}(1,2);
}

func init(){
	fmt.Println("I am first");
}