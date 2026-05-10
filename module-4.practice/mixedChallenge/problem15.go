package main;

import "fmt";

import "strconv";
type studentDB struct{
	name string;
	age int;
	grade float64;
}

func main(){
   key:="name"+ strconv.Itoa(1);
	fmt.Println(key);

	studentMap :=make(map[string]*studentDB);

	for i:=1;i<=3;i++{
		var (
			name string
			age int
			grade float64
		)
		fmt.Printf("Enter your name ");
		fmt.Scan(&name);
		fmt.Printf("Enter your age ");
		fmt.Scan(&age);
		fmt.Printf("Enter your grade ");
		fmt.Scan(&grade);
       key:=("student-"+strconv.Itoa(i));
		studentMap[key]=&studentDB{
			name:name,
			age:age,
			grade:grade,
		}

	}

	

  


	for key,value := range studentMap{
		fmt.Println("Student number is ",key);
		fmt.Println("Student Name ",value.name);
        fmt.Println("Student Age ",value.age);
		fmt.Println("Student Grade ",value.grade);
    }

	studentMap["student-1"].grade=5.00;

	//pointer dicilam na tai copy niccilo
    //    updateStudent:=studentMap["student-1"];
	//    updateStudent.grade=5.00;
    //    studentMap["student-1"]=updateStudent;
	//    fmt.Println(updateStudent);

		for key,value := range studentMap{
		fmt.Println("Student number is ",key);
		fmt.Println("Student Name ",value.name);
        fmt.Println("Student Age ",value.age);
		fmt.Println("Student Grade ",value.grade);
    }

}