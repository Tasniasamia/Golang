package main;

import "fmt";

import "strings";

func main(){

  vowelArr := []string{"A", "E", "I", "O", "U", "a", "e", "i", "o", "u"};
    var inputString string;
    fmt.Printf("Enter your string\n");
    fmt.Scan(&inputString);
    fmt.Println(inputString);


    count :=0;
    for i:=0;i<len(vowelArr);i++{
		fmt.Println(vowelArr[i]);
      conditionValue :=strings.Contains(inputString,vowelArr[i]);
      if(conditionValue){
        fmt.Println("true");
      count++;
      }

    }
    fmt.Println("The number of vowel is ",count);

    
}