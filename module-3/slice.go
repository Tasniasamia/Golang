package main;
import "fmt";

func main(){

var arr [5] int;
arr[0]=122;
arr[1]=244;


fmt.Println(len(arr));
fmt.Println(cap(arr));
fmt.Println(arr);
//way 1
arrslice:=arr[1:3];
arrslice=append(arrslice, 488);
arrslice=append(arrslice, 976);
arrslice[3]=2333;
fmt.Println("arrayslice ",arrslice);
fmt.Println("Length of arrayslice ",len(arrslice));
fmt.Println("Capacity of arrayslice ",cap(arrslice));
//way 2
arrslice2:=arr[:];
arrslice2=append(arrslice2, 488);
arrslice2=append(arrslice2, 976);
fmt.Println("arrayslice2 ",arrslice2);
fmt.Println("Length of arrayslice2 ",len(arrslice2));
fmt.Println("Capacity of arrayslice2 ",cap(arrslice2));


//slice type
newSlice:=[]int {1,2,3}
fmt.Println(newSlice);

}