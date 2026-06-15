package main;

import "fmt";



func calculate()(result int){
	result=0;
	fmt.Println("first ",result);
    show :=func(){
		result=result+10;
		fmt.Println("second ",result);
	}
	//defer-1
	defer show();

	p:=func(a int){
    fmt.Println("third ",a);
	}

	//defer-2
	defer p(result);

	//defer-3
	defer fmt.Println("forth ",result)
  

	result=5;
	fmt.Println("last function update",result);
	//defer-4
	defer fmt.Println(11);
	return
}

func calculate2()(int){
	result:=0;
	fmt.Println("first ",result);
    show :=func(){
		result=result+10;
		fmt.Println(result);
	}
	defer show();
	result=5;
	fmt.Println("last function update",result);
	return result
}







func main(){

	fmt.Println(calculate())
    //fmt.Println(calculate2());
	

}
