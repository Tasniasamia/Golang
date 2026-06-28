package main;
import "fmt";
import "time"

var a int=12;

func add(a,b int){
	fmt.Println("Custom goroutine-1 ",a+b*12);
}
func addition(a,b int){
	add(a,b)
}
func sub(a,b int){
	fmt.Println("Custom goroutine-2 ",b-a);
}

func main(){
	fmt.Println(a);

	go addition(1,2);

	go sub (2,3);

	fmt.Println(11);
	 time.Sleep(5*time.Second)



}



