// firstclasscitizenfunc.go

package main;

import "fmt";

func addMany(numbers ... int)int{
sum :=0;
for i:=0;i<len(numbers);i++{
	sum+=numbers[i];
}
return sum;
}

func mulMany(numbers ... int)int{
mul :=1;
for i:=0;i<len(numbers);i++{
	mul*=numbers[i];
}
return mul;
}

func main(){

	arrFunc := [] func(numbers ... int)(int) {addMany,mulMany}

	fmt.Println(arrFunc[0](1,2,3,4));
	fmt.Printf("%v",arrFunc); // [0x4808e0 0x480900]
	fmt.Println(arrFunc); // [0x4808e0 0x480900]
}