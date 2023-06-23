/* A struct (short for structure) is used to create a collection of members of different data types, into a single variable.
While arrays are used to store multiple values of the same data type into a single variable, structs are used to store multiple values of different data types into a single variable.
A struct can be useful for grouping data together to create records. */

package main

import (
	"fmt"
	"strings"
)

// declare a struct named "Person" which has members name, age, job and salary of diff datatypes
type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {

	// declare a variable using the struct
	var pers1 Person
	var pers2 Person

	// Pers1 specification
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000

	// Pers2 specification
	pers2.name = "Cecilie"
	pers2.age = 24
	pers2.job = "Marketing"
	pers2.salary = 4500

	// To access any member of a structure, use the dot operator (.)
	// Access and print Pers1 info
	fmt.Println("Name: ", pers1.name)
	fmt.Println("Age: ", pers1.age)
	fmt.Println("Job: ", pers1.job)
	fmt.Println("Salary: ", pers1.salary)

	// Access and print Pers2 info
	fmt.Println("Name: ", pers2.name)
	fmt.Println("Age: ", pers2.age)
	fmt.Println("Job: ", pers2.job)
	fmt.Println("Salary: ", pers2.salary)

	fmt.Println(strings.Repeat("-", 10))
	// assign value to struct while creating an instance
	pers3 := Person{"John", 25, "Actor", 80000}
	fmt.Println(pers3)

	fmt.Println(strings.Repeat("-", 10))
	// calling printPerson function defined below
	printPerson(pers1)
	fmt.Println(strings.Repeat("-", 10))
	printPerson(pers2)

}

// Pass struct to a function as argument
func printPerson(pers Person) {
	fmt.Println("printPerson function called")
	fmt.Println("Name: ", pers.name)
	fmt.Println("Age: ", pers.age)
	fmt.Println("Job: ", pers.job)
	fmt.Println("Salary: ", pers.salary)
}
