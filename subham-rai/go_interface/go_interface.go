package main

import  (
	"fmt"
	"math"
)

type  geometry  interface{
	area() float64
	perimeter() float64 
}

type circle struct {
	radius  float64
}

type rectangle struct {

	right float64
	top  float64

}

func (c circle) area() float64{
	return  math.Pi*math.Pow(c.radius,2)
}

func (c circle) perimeter() float64{
	return 2*math.Pi*c.radius
}

func (r rectangle) area() float64{
        return  r.right*r.top
}

func (r rectangle) perimeter() float64{
        return 2*r.right*r.top

}


func calculate(g geometry){
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {
	
	rect:=rectangle{right:3,top:4}
	cir:=circle{radius:3}

	calculate(rect)
	calculate(cir)
}




