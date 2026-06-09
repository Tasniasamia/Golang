package main;
import "fmt";
func call(){
	fmt.Println("call")
}
func main(){
call();
add :=func(x,y int){
		fmt.Println(x+y);
	}
	add(1,1);
	add(2,2)

}
func init(){
	fmt.Println("start");
}