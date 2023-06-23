/* The fmt.Scanf() function is similar to Scanln() function.
The only difference is that Scanf() takes inputs using format specifiers.
*/

package main

import "fmt"

func main() {
	var name string
	var age int

	// take name and age input using format specifier
	fmt.Println("Enter your name and age:")
	fmt.Scanf("%s %d", &name, &age)

	fmt.Printf("Name: %s\nAge: %d\n", name, age)

	var temperature float32
	var sunny bool

	// take float input
	fmt.Println("Enter the current temperature:")
	fmt.Scanf("%g", &temperature)

	// take boolean input
	fmt.Println("Is the day sunny?")
	fmt.Scanf("%t", &sunny)

	fmt.Printf("Current temperature: %g\nIs the day Sunny? %t", temperature, sunny)
}
