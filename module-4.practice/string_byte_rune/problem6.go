package main;

import "fmt";

func main(){
	str :="Tasnia Sharin";
	count :=0;
	for i:=0;i<len(str);i++{
		fmt.Println(str[i]);
		count++;
	}

	fmt.Println("Total count is",count);

}