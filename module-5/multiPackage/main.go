package main;

import  "multiPackage/payment";



func main(){
  //bkashPay :=bkash{apikey:"123"};
  nagodPay :=payment.Nagod{Apikey:"123"};
  makePayment :=payment.NewPaymentService(nagodPay)
  makePayment.Checkout();
	
}


