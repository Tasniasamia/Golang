package main;
import "fmt";

var a int=12;
var b int=13;

func outer() func(){
	money:=1200;
	age:=22;
	fmt.Println(age);
    show:=func(){
		money=money+a+b;
		fmt.Println(money)
	}
	return show;
}

func call(){
	incr1:=outer();
	incr1();
	incr1();

}

func main(){

	call();
}

func init(){
	fmt.Println("start");
}



















// package main;
// import "fmt";

// var a int=12;
// var b int=13;

// func outer() func(){
// 	money:=1200;
// 	age:=22;
// 	fmt.Println(age);
//     show:=func(){
// 		money=money+a+b;
// 		fmt.Println(money)
// 	}
// 	return show;
// }


// func main(){
// 	outer();
	
// }

// func init(){
// 	fmt.Println("start");
// }