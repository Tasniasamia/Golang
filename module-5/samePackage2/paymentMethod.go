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
