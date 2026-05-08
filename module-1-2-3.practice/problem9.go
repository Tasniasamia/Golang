package main;

import "fmt";

func change(x * int)(int){
	*x=13;
	return  *x;
}

func modifyArr(arr *[] int){
	for i:=0;i<len(*arr);i++{
		(*arr)[i]=(*arr)[i]*2;
	}
}

func main(){
a:=12;
fmt.Println("a is ",a);
change(&a);
fmt.Println("after change  a is ",a);

var arr[] int=[]int{1,2,3,4,5};
modifyArr(&arr);
fmt.Println(arr);

}