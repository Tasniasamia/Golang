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

type animalStruct struct{
	animal animal

}

func main(){

animalSlice:=[] animalStruct {};

meu:=cat{};
animalSlice=append(animalSlice,animalStruct{animal:meu});
pupy:=dog{};
animalSlice=append(animalSlice,animalStruct{animal:pupy});
koyel:=bird{};
animalSlice=append(animalSlice,animalStruct{animal:koyel});

fmt.Println("animalSlice",animalSlice);

for i:=0;i<3;i++{
	fmt.Println("Animal no - ",i+1);
	fmt.Printf("It's voice ");
	(animalSlice[i]).animal.speak();
	
	fmt.Printf("It eats ");
	(animalSlice[i]).animal.eat();

}




}