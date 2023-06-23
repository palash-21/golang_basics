/* We use the Scanln() function to get input values up to the new line.
When it encounters a new line, it stops taking input values.
*/

package main

import "fmt"

func main() {
	var name string
	var age int

	// take name and age input
	fmt.Println("Enter your name and age:")
	fmt.Scanln(&name, &age)

	fmt.Printf("Name: %s\nAge: %d", name, age)

}
