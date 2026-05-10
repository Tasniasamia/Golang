package main;

import "fmt";

func main(){
    var map1 map[string]int = map[string]int{};
	map1["pin"]=12;
	map1["account"]=0135;
    fmt.Println(map1);

	newMap:=map[string]int {
		"age":13,
		"roll":33,
	}
   fmt.Println(newMap);

   newMap2:=make(map[string]string);
   newMap2["name"]="Tasnia"
   fmt.Println(newMap2);

   for key,value:=range newMap2{
	fmt.Println(key,":",value);
   }


   newMap3:=[3]string {
	"hello","the","world",
   }
   fmt.Println(newMap3);

   for key,value:=range newMap3{
	fmt.Println(key,":",value);
   }

   sliceArray:=[]string {"a","b","c"};
   fmt.Printf("The type of sliceArray is %T\n", sliceArray)
   fmt.Println("sliceArray",sliceArray);

   for key,value:=range sliceArray{
	value+="1";
	fmt.Println(key,":",value);
   }
   fmt.Println("sliceArray",sliceArray);

   userName:="Tasnia Sharin";
   for _,value:=range userName{
	fmt.Printf("%c ",value);
   }
   fmt.Printf("\n");

   byteArray:=[] byte(userName);
   fmt.Println("byteArray is ",byteArray);


}