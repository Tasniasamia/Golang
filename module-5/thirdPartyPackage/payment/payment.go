package payment;
import "fmt";

type Payfunction interface{
	Pay(amount float64)
}


type paymentService struct{
 method Payfunction
}


func (p paymentService) Checkout(){
 p.method.Pay(100.0)
}
func init(){
fmt.Println("Initiating....");
}

func NewPaymentService (method Payfunction) paymentService{
	return paymentService{
		method:method,
	}
}