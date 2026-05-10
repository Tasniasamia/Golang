package main;

import "fmt";

type paymentService struct{
	thirdPartyService paymentMethod;
	amount float64;
}


func checkout(p*paymentService){

	p.thirdPartyService.pay(p.amount);
 

}
type paymentMethod interface{
	pay(amount float64)
	
}
type bkash struct{
	apikey string;
	amount float64;
}
func (b*bkash) pay(amount float64){
    fmt.Println("bkash Paid successfully",amount);
}
type nagod struct{
	apikey string;
	amount float64;
}
func (n*nagod) pay(amount float64){
  fmt.Println("nagod Paid successfully",amount);
}
type rocket struct{
	apikey string;
	amount float64;
}
func (r*rocket) pay(amount float64){
fmt.Println("rocket Paid successfully",amount);
}


func main(){
	bkashPay:=bkash{
		apikey:"123",
		amount:123.00,
	}

	paymentService1:=paymentService{
		thirdPartyService:&bkashPay,
		amount:bkashPay.amount,
	}
	checkout(&paymentService1);



	
}