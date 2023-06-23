package main

import (
	"fmt"
)

func main() {
	var i, j string = "Hello", "World"

	fmt.Print(i, j) // prints its arguments with their default format without inserting new line
	fmt.Print("\n")
	fmt.Print(i, "\n", j) // to print in a new line
	fmt.Print("\n")
	var x, y = 2, 3
	fmt.Print(x, y) //inserts a space between the arguments if neither are strings
	fmt.Print("\n")
	fmt.Print(i, j) // The Println() function is similar to Print() with the difference
	// that a whitespace is added between the arguments, and a newline is added at the end

	var p string = "Hello"
	var q int = 15

	// The Printf() function first formats its argument based on the given formatting verb and then prints them.

	// Formatting Verbs for Printf():

	// 1. General formatting verbs:
	// %v : value in default format
	// %T : datatype
	fmt.Printf("p has value: %v and type: %T\n", p, p) // print value and type of variable using Printf
	fmt.Printf("q has value: %v and type: %T\n", q, q)

	// %#v : value in Go-syntax format
	// %%  : print the % sign
	fmt.Printf("%#v\n", p)
	fmt.Printf("%v%%\n", p)
	fmt.Printf("%#v\n", q)
	fmt.Printf("%v%%\n", q)

	// 2. Integer Formatting Verbs:
	fmt.Printf("%b\n", q)   // base 2
	fmt.Printf("%d\n", q)   // base 10
	fmt.Printf("%+d\n", q)  // base 10 and always show sign
	fmt.Printf("%o\n", q)   // base 8
	fmt.Printf("%O\n", q)   // base 8 with leading 0o
	fmt.Printf("%x\n", q)   // base 16, lowercase
	fmt.Printf("%X\n", q)   // base 16, uppercase
	fmt.Printf("%#x\n", q)  // base 16, with leading 0o
	fmt.Printf("%4d\n", q)  // pad with spaces (width 4 ,right justified)
	fmt.Printf("%-4d\n", q) // pad with spaces (width 4 ,left justified)
	fmt.Printf("%04d\n", q) // pad with zeroes (width 4)

	// 3. String Formatting Verbs:
	var txt = "Hello"
	fmt.Printf("%s\n", txt)   // plain string
	fmt.Printf("%q\n", txt)   // double-quotted string
	fmt.Printf("%8s\n", txt)  // plain width 8, right justified
	fmt.Printf("%-8s\n", txt) // plain width 8, left justified
	fmt.Printf("%x\n", txt)   // hex dump of byte values
	fmt.Printf("% x\n", txt)  // 	hex dump with spaces

	// 4. Boolean Formatting Verbs:
	var b = true
	fmt.Printf("%t\n", b) // similar to %v

	// 5. Float Formatting Verbs:
	var f = 3.141
	fmt.Printf("%e\n", f)    // scientific notation with e as exponent
	fmt.Printf("%f\n", f)    // decimal point, no exponent
	fmt.Printf("%.2f\n", f)  // default width, precision 2
	fmt.Printf("%6.2f\n", f) // width 6, precision 2
	fmt.Printf("%g\n", f)    // exponent as needed, only necessary digits

	// Sprintf - to store the output of Printf in a variable
	var s string = fmt.Sprintf("%s\n", txt)
	fmt.Printf("%s\n", s) // printing value of s

}
