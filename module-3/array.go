package main;
import "fmt";

func main(){
// way 1
fmt.Println("Way 1");
var numbers [5] int;
for i:=0;i<5;i++ {
 fmt.Scan(&numbers[i]);
}
fmt.Println(numbers);


// way 2
fmt.Println("Way 2");
var arr [6] int;
arr[0]=10;
arr[2]=12;
fmt.Println(len(arr));
fmt.Println(arr);

// way 3
fmt.Println("Way 3");
var numbers2 []int =[]int{1,2,3};
fmt.Println(len(numbers2));
fmt.Println(numbers2);

numbers3 :=[]int{4,5,6};
fmt.Println(len(numbers3));
fmt.Println(numbers3);

//way 4
fmt.Println("Way 4");
var numbers4 = [6]int{100,200,300};
numbers4[3]=400;
numbers4[4]=500;
numbers4[5]=600;
fmt.Println(numbers4);

//way 5
var numbers5 [4] int;
numbers5[0]=122;
numbers5[1]=244;


fmt.Println(len(numbers5));

fmt.Println(numbers5);



//without type define array
newArray:=[3]int {1,2,3};
fmt.Println("newArray",newArray);

newArrayString:=[3]string {"a","b","c"};
fmt.Println("newArrayString",newArrayString);

}