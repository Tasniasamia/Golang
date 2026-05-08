package main;

import "fmt";
type Student struct {
 name string
 roll int
 age int
 marks [3]float64 
}

func createStudent(name string,roll int ,age int , marks [3]float64)(newStudent Student){
	newStudent=Student{name:name,roll:roll,age:age,marks:marks}
	return;
}

func avgMark(marks [3]float64)(float64){
sum:=0.00;
for i:=0;i<3;i++{
sum=sum+marks[i];
}

avg:=(sum/3);
return avg;
}

func main(){

var studentArray [] Student;
studentArraySlice:=studentArray[:]
for i:=0;i<5;i++{
 var name string
 var roll int
 var age int
 var marks [3]float64

fmt.Println("Enter your Student Info ");
fmt.Printf("Name ");
fmt.Scan(&name)
fmt.Printf("Roll ");
fmt.Scan(&roll);
fmt.Printf("Age ");
fmt.Scan(&age);
fmt.Printf("Marks ");
for i:=0;i<3;i++{
fmt.Scan(&marks[i]);
}
studentArraySlice = append(studentArraySlice, createStudent(name,roll,age,marks));
}

for i:=0;i<len(studentArraySlice);i++ {
	fmt.Printf("%d: %s | Roll: %d | Age: %d | Avg: %.2f\n",i,studentArraySlice[i].name,studentArraySlice[i].roll,studentArraySlice[i].age,avgMark(studentArraySlice[i].marks));
}

max:=avgMark(studentArraySlice[0].marks);
index:=0
for i:=1;i<len(studentArraySlice);i++{
if(avgMark(studentArraySlice[i].marks)>max){
	max=avgMark(studentArraySlice[i].marks);
	index=i;
}

}

fmt.Printf("Highest Scorer: %s (%.2f)\n",studentArraySlice[index].name,max);
} 