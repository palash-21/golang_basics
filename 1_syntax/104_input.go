// In Go, we use the scan() function to take input from the user

package main

import (
	"fmt"
	"strings"
)

func main() {
	var name string

	// takes input value for name
	fmt.Print("Enter your name: ")

	// assigns it to the name variable
	fmt.Scan(&name)

	// print assigned name
	fmt.Printf("Name: %s", name)

	// Scan() function only takes input value up to first space

	fmt.Println("")
	fmt.Println(strings.Repeat("-", 10))
	// Taking multiple inputs with Scan()
	var name1 string
	var age int

	// take name and age input
	fmt.Println("Enter your name and age:")
	fmt.Scan(&name1, &age)

	// print input values
	fmt.Printf("Name: %s\nAge: %d", name1, age)

}
