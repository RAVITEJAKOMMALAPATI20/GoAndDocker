package main

import ("fmt"
	    "math")

func main(){


	circle := Circle{5}
	rectacngle:=Rectangle{height:4,width:5} 
fmt.Println("Area of Circle of radius",getArea(circle))
fmt.Println("Area of Rectangle",getArea(rectacngle))

}





type Shape interface{
	area() float64
}
type Rectangle struct{
height float64
width float64
}


type Circle struct{
	radius float64
}
func (rect Rectangle) area() float64{
	return rect.height*rect.width
}
func (c1 Circle) area() float64{
	return math.Pi*math.Pow(c1.radius,2)
}
func  getArea(shape Shape) float64{
	return shape.area()
}