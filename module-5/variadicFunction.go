package main;

import "fmt";

func addTwo(a,b int)(int){
	return a+b;
}

func addThree(a,b,c int)(int){
	return a+b+c;
}
func addFour(a,b,c,d int)(int){
	return a+b+c+d;
}

func addMany(numbers ... int)(int){
	//numbers is number slice
	sum :=0;
	for i :=0;i<len(numbers);i++{
		sum+=numbers[i];
	}
   
	return sum;
}


func addManyWithMutipleArguments(result int,numbers ...int)(int){
	
	for i :=0;i<len(numbers);i++{
		result+=numbers[i];
	}
   
	return result;
}
func main(){
    
	//suppose we need to add two elements
	fmt.Println(addTwo(1,2));

	//we need 3 element addition
	fmt.Println(addThree(1,2,3));

	//we need 4 element addtion
	fmt.Println(addFour(1,2,3,4));

	//It's boring let's get the solution please
    fmt.Println(addMany(1,2,3,4));

	arrSlice :=[] int {10,20,30,40};
	//fmt.Println(addMany(arrSlice)); // not acceptable
	
	fmt.Println(addMany(arrSlice ...)); //acceptable 

	arrSlice=append(arrSlice,50);
	fmt.Println(addMany(arrSlice ...)); 

	//if we send two parameters what will I do let's see example
	 fmt.Println(addManyWithMutipleArguments(0,1,2,3,4,5,6,7));
     fmt.Println(addManyWithMutipleArguments(0,arrSlice ...));


	

}