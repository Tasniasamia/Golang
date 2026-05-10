package main;

import "fmt";

import "strings";

func main(){

  vowelArr := "AEIOUaeiou"
	var inputString string;
	fmt.Printf("Enter your string\n");
	fmt.Scan(&inputString);
	fmt.Println(inputString);


    count :=0;
	for i:=0;i<len(inputString);i++{
      conditionValue :=strings.Contains(vowelArr,string(inputString[i]));
	  if(conditionValue){
      count++;
	  }

	}
	fmt.Println("The number of vowel is ",count);

	
}