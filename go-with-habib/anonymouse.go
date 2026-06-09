package main;
import (
	"fmt"
)

func main(){
	a:=func(x,y int)int{
		return x+y;
	}
	fmt.Println(a(1,2));
}


