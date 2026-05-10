package main;

import "fmt";

type car struct{
	name string
}
func (c car) drive(){
	fmt.Println("Go to by ",c.name);
}
type bike struct{
	name string
}
func (b bike) drive(){
	fmt.Println("Go to by ",b.name);
}

type vehicle interface{
	drive()
}

type person struct{
	chooseVehicle vehicle; // interface type mane aitay ami ekta struct pathabo jar nijossho pay method ace
}

func main(){

	mercedes:=car{
		name:"mercedes",
	}
	person1:=person{
    chooseVehicle:&mercedes,
	}
	person1.chooseVehicle.drive();
   

}