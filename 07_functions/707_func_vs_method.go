/* A method is a function with a special receiver argument.
The receiver appears in its own argument list between the func keyword and the method name.
In this example, the Abs method has a receiver of type Vertex named v.

We will first construct a method and then implement the same as a regular function

Remember: a method is just a function with a receiver argument.
Here's Abs written as a regular function (Abs1) with no change in functionality.
*/

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {

	// Method
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	// Function
	fmt.Println(Abs1(v))
}

func Abs1(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
