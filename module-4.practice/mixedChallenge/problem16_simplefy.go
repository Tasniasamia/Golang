package main;

import "fmt";

type animal interface{
	speak()
	eat()
}

type cat struct{
}
func (c cat) speak(){
	fmt.Println("meu meu");
}
func (c cat) eat(){
	fmt.Println("fish milk")
}
type dog struct{
}
func (d dog) speak(){
	fmt.Println("gheu gheu")
}
func (d dog) eat(){
	fmt.Println("chicken milk")
}

type bird struct{
}
func (b bird) speak(){
	fmt.Println("kichir Michir")
}
func (b bird) eat(){
	fmt.Println("seeds gom");
}



func main(){

animalSlice:=[] animal {};

meu:=cat{};
animalSlice=append(animalSlice,meu);
pupy:=dog{};
animalSlice=append(animalSlice,pupy);
koyel:=bird{};
animalSlice=append(animalSlice,koyel);

fmt.Println("animalSlice",animalSlice);

for i:=0;i<3;i++{
	fmt.Println("Animal no - ",i+1);
	fmt.Printf("It's voice ");
	(animalSlice[i]).speak();
	
	fmt.Printf("It eats ");
	(animalSlice[i]).eat();

}




}