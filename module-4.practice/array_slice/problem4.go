package main;
import "fmt";
func main(){

	arr:=[]string{"Tasnia", "Tisha", "Arin"} ;

	//for range array ba slice r copy nei 
	fmt.Println("Copy of Array Or Slice")
	for _,val:=range arr {
	val+=" "+"Sharin";} 
	fmt.Println(arr);
	fmt.Printf("\n\n");


	fmt.Println("With Reference of Array Or Slice")
    for i:=0;i<len(arr);i++{
		arr[i]+=" "+"Sharin";
	}
	fmt.Println(arr);



}
