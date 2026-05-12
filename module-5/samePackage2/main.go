package main;




func main(){
  //bkashPay :=bkash{apikey:"123"};
  nagodPay :=nagod{apikey:"123"};


  
	makePayment := newPaymentService(nagodPay)
	makePayment.checkout();
	
}


//run command:  go run paymentMethod.go payment.go main.go