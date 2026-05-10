package main;

import "fmt";

func main(){
	arraysum:=[5]int{10, 20, 30, 40, 50};
	sum :=0
	for i:=0;i<5;i++{
     sum+=arraysum[i];
	}
	fmt.Println(sum);

}