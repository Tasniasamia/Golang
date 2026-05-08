package main;

import "fmt";

type bankAccount struct{
      accountHolder string;
	  accountNumber int;
	  balance float64;
	}
func createAccount(accountHolder string,accountNumber int)(newAccount bankAccount){
 newAccount=bankAccount{accountHolder:accountHolder,accountNumber:accountNumber,balance:0.00}
 return newAccount;
}
func main(){
	

	checkBalance:=func (accountInfo bankAccount) (float64) {

	
		return accountInfo.balance;
		
	}

	depositMoney:=func (amount float64,accountInfo * bankAccount) {
		
		 accountInfo.balance=accountInfo.balance+amount;
		

	}

	withdrawMoney:=func (amount float64,accountInfo * bankAccount)  {
		if(accountInfo.balance<amount){
			fmt.Println("Error: Insufficient balance!");
			return;
		}
		
		 accountInfo.balance=accountInfo.balance-amount;
		
	
 

	}



 var accountHolder string;
 var accountNumber int;


 fmt.Printf("Enter AccountHolder name ");
 fmt.Scan(&accountHolder);

 fmt.Printf("Enter your account number ");
 fmt.Scan(&accountNumber);

 var deposit float64;
 var withdraw float64
 newAccount := createAccount(accountHolder,accountNumber);

 fmt.Printf("Account created for: %s (No: %d)\n",accountHolder,accountNumber);
 fmt.Printf("Deposit ");
 fmt.Scan(&deposit);
 depositMoney(deposit,&newAccount);
 fmt.Printf("New Balance : %.2f\n",checkBalance(newAccount));
 
 fmt.Printf("Withdraw ");
 fmt.Scan(&withdraw);

 withdrawMoney(withdraw,&newAccount);

 fmt.Printf("New Balance : %.2f\n",checkBalance(newAccount));
 
 fmt.Println("Current Balance ",checkBalance(newAccount));










}