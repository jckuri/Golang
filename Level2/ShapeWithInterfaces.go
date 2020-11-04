/*
Shape with Interfaces
Statement
Create a Rectangle struct and a Circle struct, both struct should have two Method:Area() Perimeter()
And then create a Shape interface for those struct. After that create a function that will receive a Shape interface as parameter and will execute the Area() and the Perimeter() from each struct. 

Topics to Practice:
Interfaces, struct, method, function, math pkg
*/

package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (rectangle *Rectangle) Area() float64 {
	return rectangle.width * rectangle.height
}

func (rectangle *Rectangle) Perimeter() float64 {
	return 2. * (rectangle.width + rectangle.height)
}

func (circle *Circle) Area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func (circle *Circle) Perimeter() float64 {
	return 2. * math.Pi * circle.radius
}

func DescribeShape(shape Shape) string {
	return fmt.Sprintf("Type: %T, Area: %.4f, perimeter: %.4f", shape, shape.Area(), shape.Perimeter())
}

func main() {
	shapes := make([]Shape, 4)
	shapes[0] = &Rectangle{2, 5}
	shapes[1] = &Rectangle{20, 50}
	shapes[2] = &Circle{2}
	shapes[3] = &Circle{4}
	fmt.Println(shapes)
	for i := range shapes {
		fmt.Println(DescribeShape(shapes[i]))
	}
}
