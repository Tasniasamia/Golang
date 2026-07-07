package main

import "fmt"

type Users interface{
	printDetails();
	deposit(money float64)float64;
}

type user struct {
	Name string;
	Age  int;
	Amount float64;
}

func (u user) printDetails() {
fmt.Println(u.Name," ",u.Age," ",u.Amount);
}

func(u *user) deposit(money float64) float64{
u.Amount=u.Amount+money;
return u.Amount;

}

func(u *user)withdraw(money float64)float64{
u.Amount=u.Amount-money;
return u.Amount;
}


func main() {

	var u1 Users;
	u1=&user{
		Name: "Isha",
		Age:22,
		Amount:1200.00,
	}

   u1.deposit(1000.00);
   u1.printDetails();





   
	
   obj,boolean :=u1.(*user);

   if(boolean == false){
	fmt.Println("Not Valid");
	return;
   }
   obj.withdraw(500);

	u1.printDetails();

}