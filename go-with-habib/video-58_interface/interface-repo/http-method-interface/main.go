package main

import "fmt"

// import ("fmt")

type Product struct{
	Id   int    `json:"id"`
	Name string `json:"name"`
	Roll int    `json:"roll"`
	Age  int    `json:"age"`
	Dept string `json:"dept"`
}

type ProductRepo interface{

	Create(p Product)(*Product,error);
	Get(productId int)(*Product,error);
	List()([]*Product,error);
	Delete(productId int)(*Product,error);
	Update(productId int,p Product)(*Product,error);
   
}



type productRepo struct{
	productList []*Product;
}

func (r *productRepo) Create(p Product)(*Product,error){
  	if(p.Id != 0) {
		return &p,nil;
	}
	p.Id=len(r.productList)+1;
	r.productList = append(r.productList, &p)
	return &p,nil;
}

func (p *productRepo)Get(productId int)(*Product,error){
for _, p := range p.productList {
		if p.Id == productId {
			return p,nil;
		}
	}
	return nil,nil; // Return zero value if not found
}
func (p *productRepo)List()([]*Product,error){
 return p.productList,nil;
}

func (r *productRepo)Update(productId int,n Product)(*Product,error){
for _, p := range r.productList {
		if p.Id == productId {
			p= &n;
			return p,nil;
		}
	}
	return nil,nil; // Return zero value if not found
}



func (p *productRepo)Delete(productId int)(*Product,error){

	var temp*productRepo;
     var deleteProduct *Product;
	for _, p := range p.productList {
		if p.Id != productId {

			temp.productList=append(temp.productList, p)
			
		}
		deleteProduct=p;
	}
	p = temp;
	return deleteProduct,nil
}

// constructor or constructor function

func createProductRepo () ProductRepo {

  repo := &productRepo{};

  generateRepo(repo);

  return repo;

}












func generateRepo(s*productRepo){
	u1 := &Product{
		Id:   1,
		Name: "John Doe",
		Roll: 101,
		Age:  25,
		Dept: "Computer Science",
	}
	u2 := &Product{
		Id:   2,
		Name: "Jane Smith",
		Roll: 102,
		Age:  23,

		Dept: "Mathematics",
	}
	u3 := &Product{
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






// func (u product) Store() User {
// 	if(u.Id != 0) {
// 		return u;
// 	}
// 	u.Id=len(users)+1;
// 	users = append(users, u)
// 	return u;
// }

// func Find(id int) *User {
// 	for _, user := range users {
// 		if user.Id == id {
// 			return &user
// 		}
// 	}
// 	return nil // Return zero value if not found
// }

// func Update(id int, updatedUser User) User {
// 	for i, user := range users {
// 		if user.Id == id {
// 			users[i] = updatedUser
// 			return updatedUser
// 		}
// 	}
// 	return User{} // Return zero value if not found
// }

// func Delete(id int) {
// 	var temp [] User;
// 	for _, user := range users {
// 		if user.Id != id {

// 			temp=append(temp, user)
			
// 		}
// 	}
// 	users = temp;
// }