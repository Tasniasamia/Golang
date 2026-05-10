package main;
import "fmt";



// range করলে → rune (int32), 4 byte, Unicode safe
// []byte করলে → byte (uint8), 1 byte, raw memory

func main(){
 newString:="Hello";
  for _,val := range newString {
	fmt.Printf("%T %d ",val,val);
  }
 
  fmt.Printf("\n\n")
  stringToByte:=[]byte(newString);
  fmt.Printf("%T %v",stringToByte,stringToByte);

    fmt.Printf("\n\n")
    for _,val := range stringToByte {
	fmt.Printf("%T %d ",val,val);
  }
}