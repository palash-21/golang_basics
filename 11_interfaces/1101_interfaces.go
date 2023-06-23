/* Go programming provides another data type called interfaces which represents a set of method signatures
or in other words named collections of method signatures.
The struct data type implements these interfaces to have method definitions for the method signature of the interfaces.
We use interfaces to store a set of methods without implementation. That is, methods of interface won't have a method body.
*/

package main

import (
	"fmt"
	"math"
)

/* define an interface */
type Shape interface {
	area() float64
}

/* define a circle */
type Circle struct {
	x, y, radius float64
}

/* define a rectangle */
type Rectangle struct {
	width, height float64
}

/* define a method for circle (implementation of Shape.area())*/
func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

/* define a method for rectangle (implementation of Shape.area())*/
func (rect Rectangle) area() float64 {
	return rect.width * rect.height
}

/* define a method for shape */
func getArea(shape Shape) float64 {
	return shape.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{width: 10, height: 5}

	fmt.Printf("Circle area: %f\n", getArea(circle))
	fmt.Printf("Rectangle area: %f\n", getArea(rectangle))
}
