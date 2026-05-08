package main
import "fmt"


func main(){
 
getData:=func (){
	fmt.Println("Hello world");

}

getData();

func(){
	fmt.Println("I will never forgive my mother");
}()

func(){
fmt.Println("My mother was never strong");
}()

func(){
	fmt.Println("She was always depends on one person");
}()


//if else scope

score:=89;
fmt.Println(score);


if score:=70;score>=80 {
	fmt.Println("A+");
}else if ((score>=70) && (score<80)){
	fmt.Println("A");
}else{
	fmt.Println("C");
}


}