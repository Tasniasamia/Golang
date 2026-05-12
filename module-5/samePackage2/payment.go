package main;

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
		method:method,
	}
}