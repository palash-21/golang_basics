/* 
A Go file consists of the following parts:
	Package declaration
	Import packages
	Functions
	Statements and expressions
	Variables
	Comments
*/

package main    // Everything in go program is written inside a package

import (
	"fmt"       // import format module
	"math/rand" // import random from math module
)

func main() {				     // to use this (main) function, we must import the package main
	fmt.Println("Hello World!")  // Print line function prints the output and inserts a new line
	fmt.Println("My favorite number is", rand.Intn(10))  // print a random integer in range [0,n) 
}

// to run this file use cmd : 'go run filename.go' ---> function main will execute
// to build it use 'go build filename.go' ---> generate a .exe file that can run directly without any dependencies.

/* Notice the capital P of Println method. In Go language, a name is exported if it starts with capital letter. 
 Exported means the function or variable/constant is accessible to the importer of the respective package */

// Above code in compact syntax
// package main; import ("fmt"); func main() { fmt.Println("Hello World!");}


/* By convention, the package name is the same as the last element of the import path. 
 For instance, the "math/rand" package comprises files that begin with the statement package rand.
 When a package is imported, only entities (functions, types, variables, constants) whose names start with a capital letter can be used / accessed. 
 The recommended style of naming in Go is that identifiers will be named using camelCase, 
 except for those meant to be accessible across packages which should be PascalCase. 
 eg. math.pi will panic but math.Pi will access the constant Pi */

