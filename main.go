package main
import "fmt"



//normal function
func cook(dish string, isSalty bool){
	fmt.Printf("This dish is %s\n",dish);
	fmt.Printf("IsSalty %t\n",isSalty);
}

//multi value return function
func pizza(item string, length int)(string,int){
	made:=fmt.Sprintf("%s Pizza and its length %d",item,length);
	return made,length
}
func travel(place string, outofCountry bool)(foreign bool, tourPlace string){
	tourPlace=place;
	foreign=outofCountry;
	return ;
}
func main(){
	
	//Hello World Program
	fmt.Println("Hello World");
	
	//Variable
	var name string="Tasnia";
	fmt.Println(name)
	age:=18
	fmt.Println(age)

	var(
		department string="computer"
		section string="4P"
		year int=4
	);
	fmt.Println(department);
	fmt.Println(section);
	fmt.Println(year);


//multi variable
fmt.Println("Multi Variable");
var x,y int =3,4 ;
// x=1;
// y=2;
fmt.Println(x);
fmt.Println(y);

var userName string="disha";
fmt.Println(userName);

const roll=12;
fmt.Println(roll);


//zero values and printf
fmt.Println("zero values and printf");
fmt.Printf("The variable of x is %d and y is %d \n",x,y);

user:=fmt.Sprintf("The variable of x is %d ",x);
fmt.Println(user);

//normal function
cook("Pizza",true);

//multiple value return function  
pizzaName,pizzaInch:= pizza("chicken",8);
fmt.Println(pizzaName);
fmt.Println(pizzaInch);

foreign,tourPlace:=travel("Maldives",true);
fmt.Printf("Tour place is %s\n",tourPlace);
fmt.Printf("Is Foreign %t\n",foreign);



}