package main;
import "fmt"
func main(){
    for(true){
	fmt.Println("Menu")
	func(){
		fmt.Println("1. Calculate Grage");
		fmt.Println("2. Check Pass/Fail");
		fmt.Println("3. Exit");
    }()

	var choose int;
	var score int;
	fmt.Printf("Enter your choice ")
	fmt.Scan(&choose);

    switch (choose){
    case 1:
	fmt.Printf("Enter Score ");
	fmt.Scan(&score);
	if (score>=90) && (score<=100){
    fmt.Println("A");
	}else if (score>=80) && (score<=89){
     fmt.Println("B");
	} else if (score>=70) && (score<=79){
    fmt.Println("C");
	}else if (score>=60) && (score<=69){
    fmt.Println("D");
	} else{
	fmt.Println("F");
    }
	case 2:
	fmt.Printf("Enter Score ");
	fmt.Scan(&score);
	
	conditon:=score>=60;
	switch(conditon){
	case true:
	fmt.Println("Pass");
	case false:
	fmt.Println("Fail");
	}
	case 3:
    break;
   }
   if(choose == 3){
	break;
   }
}
}