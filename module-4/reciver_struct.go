package main;

import "fmt";

type user struct{
	name string
	age int
	isLoggedIn bool
}
func main(){

	user1:=user{
		name:"Tasnia",
		age:24,
	};

	
	user1.login(true);
	user1.updateInfo();
	fmt.Println("User1 is ",user1);
	
	user2:=user{
		name:"Tisha",
		age:26,
	};
	user2.updateInfo();
    user2.login(true);
	fmt.Println("User2 is ",user2);


}

func (u*user) updateInfo (){
	u.name="Arin";
	u.age=22;
}




func (u*user) login(crediential bool){
if(crediential == true){
 u.isLoggedIn=true;
 fmt.Println("Logged in");
 return;
}
 u.isLoggedIn=false;
 fmt.Println("Not Logged in");
}