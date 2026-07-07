package main

import "fmt"

// import ("fmt")

type product struct{
	Id   int    `json:"id"`
	Name string `json:"name"`
	Roll int    `json:"roll"`
	Age  int    `json:"age"`
	Dept string `json:"dept"`
}

type ProductRepo interface{

	// print();

}

type productRepo struct{
	productList []product;
}

// func(s*productRepo)print(){
//   fmt.Println(s);
// }
func createProductRepo () ProductRepo {

  repo := &productRepo{productList:[]product{}};

  generateRepo(repo);

  return repo

}

func generateRepo(s*productRepo){
	u1 := product{
		Id:   1,
		Name: "John Doe",
		Roll: 101,
		Age:  25,
		Dept: "Computer Science",
	}
	u2 := product{
		Id:   2,
		Name: "Jane Smith",
		Roll: 102,
		Age:  23,

		Dept: "Mathematics",
	}
	u3 := product{
		Id:   3,
		Name: "Bob Johnson",
		Roll: 103,
		Age:  24,
		Dept: "Physics",
	}
	
	s.productList = append(s.productList, u1,u2,u3);
}


func main(){
	s:=createProductRepo();
	fmt.Println(s);
}