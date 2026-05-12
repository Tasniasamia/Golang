package main;

import "thirdPartyPackage/methods";

import "thirdPartyPackage/payment";


func main(){
  //bkashPay :=bkash{apikey:"123"};
  nagodPay :=methods.Nagod{Apikey:"123"};


  
	makePayment := payment.NewPaymentService(nagodPay)
	makePayment.Checkout();
	
}


//run command:  go run paymentMethod.go payment.go main.go