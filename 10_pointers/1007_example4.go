// Program to access the field of a struct using pointer

package main

import "fmt"

func main() {

	// declare a struct Person
	type Person struct {
		name string
		age  int
	}

	person := Person{"John", 25}

	// create a struct type pointer that
	// stores the address of person
	var ptr *Person
	ptr = &person

	// access the name member using pointer
	fmt.Println("Name:", ptr.name)

	// access the age member using pointer
	fmt.Println("Age:", ptr.age)

	// change the member value using pointer
	ptr.age = 20
	fmt.Println(person)

}
