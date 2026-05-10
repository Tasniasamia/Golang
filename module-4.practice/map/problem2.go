package main;
import "fmt";
import "strings";

func main(){
	stringArr := "go is fun and go is easy";
	newStrginArr := strings.Split(stringArr," ");
	fmt.Println(newStrginArr);

	newStringMap:=make(map[string]int);

	for _,value:= range newStrginArr{
		_,ok:=newStringMap[value];

		if(ok ){
			newStringMap[value]= newStringMap[value]+1;
		}else{
		newStringMap[value]=1;

		}
	}

	fmt.Println(newStringMap);
}