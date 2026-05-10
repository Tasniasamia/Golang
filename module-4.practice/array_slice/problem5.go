package main;

import "fmt";

func main(){
	arr :=[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
	newSlice :=[]int{};
	for i:=0;i<len(arr);i++{
		modulus:=(arr[i]%2);
		if(modulus==0){
		newSlice=append(newSlice,arr[i]);
		}
		
	}
	fmt.Println(newSlice);
}