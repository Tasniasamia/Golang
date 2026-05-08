package main;
import "fmt";

func coffeeOrder(name string,coffee string,price int)(string){
	output:=fmt.Sprintf("Order for %s: %s coffee costs %d taka ☕",name,coffee,price);
	return output;
}

func main(){
order := coffeeOrder("Mizan", "Latte", 150)
fmt.Println(order)
}

