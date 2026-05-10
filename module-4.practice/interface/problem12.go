package main;

import "fmt";

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

func checkout(p paymentMethod,amount float64){

	p.pay(amount);

}
func main(){
bkashPay:=bkash{
	apikey:"123456",
	amount:1200.00,
}
nagodPay:=nagod{
	apikey:"123456",
	amount:1200.00,
}
rocketPay:=rocket{
	apikey:"123456",
	amount:1200.00,
}
checkout(&bkashPay,bkashPay.amount);
checkout(&nagodPay,nagodPay.amount);
checkout(&rocketPay,rocketPay.amount);


}