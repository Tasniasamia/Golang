package main;
import "fmt";

func main(){

	//problem 1
	fmt.Println("Problem 1");
    Name:="Tasnia"
	Age:=23;
	GPA:=3.88
	const PI float32=3.1416;
	fmt.Printf("Name:%s\n",Name);
	fmt.Printf("Age:%d\n",Age);
	fmt.Printf("GPA:%.2f\n",GPA);
	fmt.Printf("PI:%.4f\n\n",PI);

//problem 2
fmt.Println("Problem 2");
var integer int;
var floating float64;
var stringValue string;
var boolean bool;
integer=0;
floating=0;
stringValue="";
boolean=false;
fmt.Printf("int zero value: %d (type %T)\n",integer,integer);
fmt.Printf("floating zero value: %f (type %T)\n",floating,floating);
fmt.Printf("stringValue zero value: %s (type %T)\n",stringValue,stringValue);
fmt.Printf("boolean zero value: %t (type %T)\n\n",boolean,boolean);


//problem 3
fmt.Println("Problem 3");
fmt.Printf("sum is %d\n",sum(1,2));
fmt.Printf("Number is %s\n",checkEvenOdd((6)));
area,perimerter :=rectangle(2,4);
fmt.Printf("Rectange area is %.2f and perimeter is %.2f\n\n",area,perimerter);



//problem 4
fmt.Println("Problem 4");
var age int;
fmt.Printf("Enter your age ")
fmt.Scan(&age);
if(age<18){
	fmt.Println("Minor");
} else if((age>=18) && (age<=60)){
	fmt.Println("Adult")
}else if(age>60){
	fmt.Println("Senior");
}else{
	fmt.Println("Invalid");
}

var weeKDay_number int;
fmt.Printf("Enter your Weekday Number ")
fmt.Scan(&weeKDay_number);
switch(weeKDay_number){
case 1:
fmt.Println("Saturday");
case 2:
fmt.Println("Sunday");
case 3:
fmt.Println("Monday");
case 4:
fmt.Println("Tuesday");
case 5:
fmt.Println("Wednesday");
case 6:
fmt.Println("Thursday");
case 7:
fmt.Println("Friday");
default:
fmt.Println("Invalid");
}

switch score:=80; {
case (score>=80):
	fmt.Println("A");
case (score>=60):
	fmt.Printf("B");
case (score>=40):
	fmt.Println("C");
default:
	fmt.Println("Fail");
}
fmt.Printf("\n")

//problem-5 
fmt.Println("Problem 5");
sum:=0;
for i:=0;i<100;i++{
sum=sum+i;
}
fmt.Println("The sum is ",sum);

var num int;
fmt.Printf("Enter the number to multiply ");
fmt.Scan(&num);
for i:=1;i<=10;i++{
fmt.Printf("%d * %d = %d\n",num,i,(num*i));
}

for i:=1;i<=50;i++{
	if(i%2==0){
		continue;
	}
	fmt.Printf("Odd number is %d\n",i);
}

fmt.Printf("\n");


//Problem 6
fmt.Println("Problem 6");
var arr [5] int ;
for i:=0;i<5;i++{
	fmt.Scan(&arr[i]);
}
fmt.Println(arr);
min:=arr[0];
max:=arr[0];
sum=arr[0];
for i:=1;i<5;i++{
	sum=sum+arr[i];
	if(arr[i]<min){
		min=arr[i];
	}
	if(arr[i]>max){
		max=arr[i];
	}
}

fmt.Printf("Minimum value is %d\n",min);
fmt.Printf("Maximum value is %d\n",max);
fmt.Printf("Average Value is %d\n",sum/5);
}

//problem 3
func sum (a int,b int)(int){
	return a+b;
}

func checkEvenOdd(num int)(string){
	if(num%2==0){
		return "Even";
	} else {
		return "Odd";
	}
}

func rectangle(length float64,width float64)(area float64,perimeter float64){
area = (length*width);
perimeter =(2.0*(length+width));
return ;
}