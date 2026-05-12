package methods;
import "fmt";

type bkash struct{
apikey string;
}

func (b bkash) Pay(amount float64){
fmt.Println("ay ",amount);
}

type Nagod struct{
Apikey string;
}

func (n Nagod) Pay(amount float64){
fmt.Println("pay ",amount);
}
