package main;

import "fmt"

type bankAccount struct{
	owner string;
	balance float64;
}

func (b*bankAccount) deposit(amount float64){
	b.balance+=amount;
}

func (b*bankAccount) withdraw(amount float64){
	if(b.balance<amount){
		fmt.Println("Insufficient Balance");
		return;
	}
	b.balance-=amount;
}

func (b*bankAccount) getBalance(){
	fmt.Println(b.balance);
}
func main(){

	newAccount :=bankAccount{
		owner:"Disha",
		balance:0.00,
	}

newAccount.getBalance();
newAccount.deposit(12000.000);
newAccount.withdraw(13000.000);
newAccount.getBalance();

}