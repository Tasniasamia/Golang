package main;

import "fmt"

func main(){
 type user struct{
	name string;
	roll int;
	age int;
	college string;
 }
 person1:=user{};
 //var person1 user //same
 person1.name="samia";
 person1.roll=12;
 fmt.Println(person1.name);
 fmt.Println(person1.roll);
 fmt.Println("struct person1 ",person1);


 studentArray:=[5]user{};
// var studentArray2 user;  // same
 for i:=0;i<5;i++ {

fmt.Print("Name is: ")
fmt.Scan(&studentArray[i].name);

fmt.Print("Roll is: ")
fmt.Scan(&studentArray[i].roll)

fmt.Print("Age is: ")
fmt.Scan(&studentArray[i].age);

fmt.Print("College is: ")
fmt.Scan(&studentArray[i].college);
 }




 for i:=0;i<5;i++ {
fmt.Printf("Person-%d\n",i+1);
fmt.Printf("Name is: %s\n",studentArray[i].name);
fmt.Printf("Roll is: %d\n",studentArray[i].roll);
fmt.Printf("Age is: %d\n",studentArray[i].age);
fmt.Printf("College is: %s\n",studentArray[i].college);



 }


}
