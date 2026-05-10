package main;

import "fmt"

type rectange struct{
	width float64;
	height float64;
}

func (r*rectange) area()(float64){
  return (r.width*r.height);
}

func (r*rectange) perimeter()(float64){
  return (2*(r.width+r.height));
}

func (r*rectange) isSqure()(bool){
  return (r.width==r.height);
}

func main(){

	rectangle1:=rectange{
		width:4.0,
		height:4.0,
	}
	fmt.Println(rectangle1.area())
	fmt.Println(rectangle1.perimeter())
	fmt.Println(rectangle1.isSqure())


}