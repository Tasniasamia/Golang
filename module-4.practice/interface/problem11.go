package main;

import "fmt";

type shape interface{
	area()float64
	perimeter()float64
}

type rectange struct{
	height float64;
	width float64;
}

type circle struct{
	radius float64
}

func (r*rectange) area()float64{
	return r.height*r.width;
}

func (r*rectange) perimeter()float64{
	return (2*(r.height+r.width));
}

func (c*circle) area()float64{
	return 3.1416*(c.radius*c.radius);
}
func (c*circle) perimeter()float64{
	return 2*3.1416*c.radius;
}

func shapeInfo(s shape){

	fmt.Println("Area is",s.area());
	fmt.Println("Perimeter is",s.perimeter());

}

func main(){
rectange1 :=rectange{
	height:4.0,
	width:4.0,

}

circle1:=circle{
	radius:4.5,
}
shapeInfo(&rectange1);
shapeInfo(&circle1);

}