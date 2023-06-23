package main

import (
	"fmt"
)

// simple function
func myMessage() {
	fmt.Println("I just got executed!")
}

// function with parameters
func familyName(fname string) { // function with parameter fname of type string
	fmt.Println("Hello", fname, "Ryan")
}

// function with muli parameters
func familyNameWithAge(fname string, age int) {
	fmt.Println("Hello", age, "year old", fname, "Refsnes")
}

// function with return
func myFunction(x int, y int) int { // function takes two integers and return integer
	return x + y
}

// function with named return
func myFunction1(x int, y int) (result int) {
	result = x + y
	return // naked return
}

// function with multi return value
func myFunction3(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

func main() { // main function gets executed with this program , main func doesnot accept parameters/ or return anything

	myMessage() // call the function

	familyName("Jack") // pass firstname as arguments to the function to print full name
	familyName("John")

	familyNameWithAge("Jack", 34) //When you are working with multiple parameters, the function call must have the same number of arguments as there are parameters,
	// and the arguments must be passed in the same order.

	fmt.Println(myFunction(5, 10))
	fmt.Println(myFunction1(5, 10))

	a, b := myFunction3(5, "Hello") // store the returns values in variables
	fmt.Println(a, b)

}
