package methods;
// import "fmt";
import "github.com/fatih/color"

type bkash struct{
apikey string;
}

func (b bkash) Pay(amount float64){
color.Red("Paid Amount %v\n",amount)
}

type Nagod struct{
Apikey string;
}

func (n Nagod) Pay(amount float64){
color.Red("Paid Amount %f\n",amount)
}
