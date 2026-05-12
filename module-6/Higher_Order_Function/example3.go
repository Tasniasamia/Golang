package main;

import "fmt";



func callBack(a func(b func(c func()),d func()),x func(m func()),y func()){
	
	a(x,y);
	
	
}

func func1(){
 fmt.Println("last functions");
}

func func2(func4 func()){

	 func4();

}

func func3(func22 func(funcNew func()),afunc func()){

	 func22(afunc);
}


func main(){



callBack(func3,func2,func1);
	
	
}