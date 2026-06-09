// package main;
// import "fmt"



// type dog struct {
//    name string
// }

// type cat struct {
//    name string
// }
// type common interface{
// 	voice()
// }
// func (c cat) voice() {
//    fmt.Println("Meu")
// }

// func (d dog) voice() {
//    fmt.Println("Gheu")
// }

// func makeSound(c common){
//  c.voice();
// }

// func main() {
//    cat1 := cat{
//        name: "meu",
//    }

//    dog1 := dog{
//        name: "pupy",
//    }

//    makeSound(cat1)   // Meu
//    makeSound(dog1)
  
// }


package main;
import "fmt";

func main(){
var name int;
fmt.Printf("Enter number of a is ");
fmt.Scan(&name);
if (name==12) {
	fmt.Println(name);
}

arr :=[]int{1,2,3,4};
fmt.Printf("%v\n",arr);
}