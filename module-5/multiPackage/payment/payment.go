package payment;
import "fmt"

type bkash struct{
apikey string;
}

func (b bkash) pay(amount float64){
fmt.Println("pay ",amount);
}

type Nagod struct{
Apikey string;
}

func (n Nagod) pay(amount float64){
fmt.Println("pay ",amount);
}



type Payfunction interface{
	pay(amount float64)
}


type paymentService struct{
 method Payfunction
}


func (p paymentService) Checkout(){
 p.method.pay(100.0)
}


func NewPaymentService (method Payfunction) paymentService{
	return paymentService{
		method:method,
	}
}