/* Identifers
A Go identifier is a name used to identify a variable, function, or any other user-defined item
 An identifier starts with a letter A to Z or a to z or an underscore _
 followed by zero or more letters, underscores, and digits (0 to 9).

Go does not allow punctuation characters such as @, $, and % within identifiers.
A variable name cannot contain spaces.
One should avoid using go keywords eg. break, func, map, case, etc
*/

// Declaring variables
// Either type or value or both must be provided to declare a variable

package main

import (
	"fmt"
)

func main() {
	var student1 string = "John" //type is string
	var student2 = "Jane"        //type is inferred
	x := 2                       //type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

	var a, b, c, d int = 1, 3, 5, 7 // declare multiple variable of same type
	var e, f = 6, "Hello"           // declare multiple variable of diff type
	fmt.Println(a, b, c, d, e, f)

	var ( // declare multiple variable of diff type in a block
		i int
		j int    = 1
		k string = "hello"
	)
	fmt.Println(i, j, k)
	const PI = 3.14 // declaring a constant, standard practice to use capital letters
	// we cannot use the shorthand notation := to create constants
}

/* In Go, all variables are initialized.
So, if you declare a variable without an initial value, its value will be set to the default value of its type
eg. default value of int is 0, string is '' and so on
So, deckaring a varibale x of type int (var x int ) will set x value to 0
*/

// Note that ':=' can only be used inside a function
