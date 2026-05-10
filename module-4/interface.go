//without interface

// package main;

// import "fmt";

// type cat struct{
// 	name string;
// 	age int
// }
// type dog struct{
// 	name string;
// 	age int;
// }
// func (c*cat) speak(){
// 	fmt.Println("Meu Meu");
// }
// func (c*cat) eat(){
// 	fmt.Println("Fish");
// }
// func (d*dog) speak(){
// 	fmt.Println("Gheu Gheu");
// }
// func makeSound(c*cat){
// c.speak();
// }
// func makeSound2(d*dog){
// d.speak();
// }

// func main(){
	
// 	meu :=cat{
// 		name:"meu",
// 		age:1,
// 	}
// 	makeSound(&meu);
   
//    puppy :=dog{
// 		name:"puppy",
// 		age:2,
// 	}
// 	makeSound2(&puppy);
// 	fmt.Println(puppy);
// }


//with interface


package main;

import "fmt";

type animal interface{
	speak();
}

func makeSound(a animal){
a.speak();
}

type cat struct{
	name string;
	age int
}
type dog struct{
	name string;
	age int;
}

func (c*cat) speak(){
	fmt.Println("Meu Meu");
}

func (c*cat) eat(){
	fmt.Println("Fish");
}

func (d*dog) speak(){
	fmt.Println("Gheu Gheu");
}



func main(){
	
	meu :=cat{
		name:"meu",
		age:1,
	}
	//pointerMeu:=&meu;
	//makeSound(pointerMeu);
    makeSound(&meu);


	puppy :=dog{
		name:"puppy",
		age:2,
	}
	fmt.Println(puppy);

}