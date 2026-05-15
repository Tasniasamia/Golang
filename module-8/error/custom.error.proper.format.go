package main;

import (
	"fmt"
	// "errors"
)

type customError struct{
	message string;
	code     int;
	success  bool;
}

func (c*customError) Error()string{
	return c.message;
}


func wrongPassword(password string)error{
  return &customError{
	message:"Wrong password here",
	code:401,
	success:false,
  }
}

func main(){
	

err :=wrongPassword("1234");

if result,ok :=err.(*customError);ok{
fmt.Printf("status code:%d success:%t message:%s\n",result.code,result.success,result.message);
}else{
fmt.Println("Something went wrong");

}


}