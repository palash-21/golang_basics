// Go does not hvae in built method support but we can define a method using functions

package main

import (
	"fmt"
	"math"
)

/* define a circle */
// struct is a collection of variables (of one/more datatypes)
type Circle struct {
	x, y, radius float64
}

/* define a method for circle */
func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	fmt.Printf("Circle area: %f", circle.area())
}
