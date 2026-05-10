package main;

import "fmt";

type bkash struct{
apikey string;
}

func (b bkash) pay(amount float64){
fmt.Println("pay ",amount);
}

type nagod struct{
apikey string;
}

func (n nagod) pay(amount float64){
fmt.Println("pay ",amount);
}



type Payfunction interface{
	pay(amount float64)
}


type paymentService struct{
 method Payfunction
}


func (p paymentService) checkout(){
 p.method.pay(100.0)
}


func newPaymentService (method Payfunction) paymentService{
	return paymentService{
		method:method
	}
}

func main(){
  //bkashPay :=bkash{apikey:"123"};
  nagodPay :=nagod{apikey:"123"};


  
	makePayment := newPaymentService(nagodPay)
	makePayment.checkout();
	
}