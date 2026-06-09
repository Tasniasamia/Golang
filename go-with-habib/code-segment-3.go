package main;
import "fmt";
func call(){
	fmt.Println("call")
}
func main(){
call();
add :=func(x,y int)int{
		return x+y;
	}(1,2);

	fmt.Println(add);
	
}
func init(){
	fmt.Println("start");
}