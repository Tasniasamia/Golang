package main;

import "fmt";

func main(){
	arr:=[]int{};
	
	arraySlice:=arr[:];
	arraySlice = append(arraySlice, 1);
	arraySlice = append(arraySlice, 2);
	arraySlice = append(arraySlice, 3);
	arraySlice = append(arraySlice, 4);
	arraySlice = append(arraySlice, 5);

	fmt.Println(arraySlice);
	fmt.Println("Length of arraySlice is ",len(arraySlice))
	fmt.Println("Capacity of arraySlice is ",cap(arraySlice))
	newSlice:=arraySlice[1:4];
	fmt.Println("Index 1 to 3 of slice is ",newSlice);

	

}