package main
// linke ref https://ithelp.ithome.com.tw/articles/10204662
// and https://medium.com/golang-%E7%AD%86%E8%A8%98/golang-interface-oo-note-14fb1cb76600
import (
	"fmt"
	"math"
)

// Define interface

// Shape --  area method
type Shape interface {
	area() float64
	
}
// Circle -- 
type Circle struct {
	x, y, radius float64
}
// Rectangle --
type Rectangle struct {
	width, height float64
}

func (c Circle) halfArea() float64 {
	return math.Pi * c.radius * c.radius /2
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{width: 10, height: 5}

	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Rectangle Area: %f\n", getArea(rectangle))
}