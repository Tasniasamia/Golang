package main;

import "fmt";

type user struct{
	name string;
	roll int;
}

func main(){
	//empty map create way1
	var map2 map[string]int=map[string]int {} ;
	fmt.Println(map2);

     // empty map create way2
     newMap:=make(map[string]int)
     newMap["age"]=12;
	 fmt.Println(newMap);

	 //map create no empty
	 newMap3:=map[string]int {
		"age":12,
		"roll":1,
	 }
	 newMap3["id"]=1234;
	 fmt.Println(newMap3);

     //map key delete
	 delete(newMap3,"id");
	fmt.Println(newMap3);


	//struct map create 
	structMap:= map[string]user {
		"user1":user{
			name:"Tasnia",
			roll:1,
		},
		"user2":user{
			name:"Tina",
			roll:2,
		},
	}

	fmt.Println(structMap);


}