package main;

import (
	"errors"
	"fmt"
)

func doSomething(a,b int)(int,error){
  if((a==0) || (b==0)){
   return 0,errors.New("0 is not allowed as parameter")
  }

  return a*b,nil;
}

func main(){
	result,err :=doSomething(2,0);
	if(err != nil){
		fmt.Println(err);
		return
	
	}
   fmt.Println(result);
}