package main;
import (
	"fmt"
)

var num int=12;

func main(){
	func(x,y int){
		fmt.Println(x+y);
	}(1,2);

	fmt.Println(num);
	
}


