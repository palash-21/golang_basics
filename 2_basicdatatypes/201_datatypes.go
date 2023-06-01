/* The type of a variable determines how much space it occupies in storage, how the bit pattern stored is interpreted,
   set of operations that can be applied on the variable */

package main

import (
	"fmt"
	"strings"
)

func main() {

	// Boolean - true or false
	// Default = false
	var b1 bool = true // typed declaration with initial value
	var b2 = true      // untyped declaration with initial value
	var b3 bool        // typed declaration without initial value
	b4 := true         // untyped declaration with initial value
	fmt.Println(b1)    // Returns true
	fmt.Println(b2)    // Returns true
	fmt.Println(b3)    // Returns false
	fmt.Println(b4)    // Returns true

	// Integer
	// Default = 0
	// Default type of integer is int if not specified

	// int is signed type integer datatype (both positive and negative values)
	var x int = 500
	var y int = -4500
	fmt.Printf("Type: %T, value: %v\n", x, x)
	fmt.Printf("Type: %T, value: %v\n", y, y)

	// int type with sizes
	/*
		int   : 32bit or 64bit depending on O.S.
		int8  : 8  bits/ 1 byte (-128 to 127)
		int16 : 16 bits/ 2 byte
		int32 : 32 bits/ 4 byte
		int64 : 64 buts/ 8 byte
	*/

	// uint is unsigned integer datatues (only positive values and zero)
	var u uint = 500
	fmt.Printf("Type: %T, value: %v\n", u, u)

	// uint type with sizes
	/*
		uint   : 32bit or 64bit depending on O.S.
		uint8  : 8  bits/ 1 byte (0 to 255))
		uint16 : 16 bits/ 2 byte
		uint32 : 32 bits/ 4 byte
		uint64 : 64 buts/ 8 byte
	*/

	// Float
	// Default = 0
	// Default type of float is float64 if not specified

	// float32 : 32 bits and float64 : 64 bits

	var f1 float32 = 3.4e+38
	fmt.Printf("Type: %T, value: %v\n", f1, f1)

	var f2 float64 = 1.7e+308 // float64 can store decimals with greater precision
	fmt.Printf("Type: %T, value: %v\n", f2, f2)

	// String
	// Default ""

	var txt1 string = "Hello!"
	var txt2 string
	txt3 := "World 1"

	fmt.Printf("Type: %T, value: %v\n", txt1, txt1)
	fmt.Printf("Type: %T, value: %v\n", txt2, txt2)
	fmt.Printf("Type: %T, value: %v\n", txt3, txt3)

	/*Strings, which are widely used in Go programming, are a readonly slice of bytes.
	In the Go programming language, strings are slices. The Go platform provides various libraries to manipulate strings.

	unicode
	regexp
	strings
	*/

	var greeting = "Hello world!"

	fmt.Printf("normal string: ")
	fmt.Printf("%s", greeting)
	fmt.Printf("\n")
	fmt.Printf("hex bytes: ")

	for i := 0; i < len(greeting); i++ {
		fmt.Printf("%x ", greeting[i])
	}

	fmt.Printf("\n")
	const sampleText = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	/*q flag escapes unprintable characters, with + flag it escapses non-ascii
	characters as well to make output unambigous */
	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", sampleText)
	fmt.Printf("\n")

	// length of string function
	fmt.Println(len(greeting))

	// concate with +
	message1 := "I love"
	message2 := "Go programming"

	// concatenation using + operator
	result := message1 + " " + message2

	fmt.Println(result)

	// The strings package includes a method join for concatenating multiple strings
	greetings := []string{"Hello", "world!"}
	fmt.Println(strings.Join(greetings, " "))

	// compare strings
	string1 := "Programiz"
	string2 := "Programiz Pro"
	string3 := "Programiz"

	// compare strings
	fmt.Println(strings.Compare(string1, string2)) // -1 because string1 is smaller than string2
	fmt.Println(strings.Compare(string2, string3)) // 1  because string3 is smaller than string2
	fmt.Println(strings.Compare(string1, string3)) // 0  because string1 is equal to string3

	// containes
	text := "Go Programming"
	substring1 := "Go"
	substring2 := "Golang"

	// check if Go is present in Go Programming
	result_contains := strings.Contains(text, substring1)
	fmt.Println(result_contains)

	// check if Golang is present in Go Programming
	result_contains = strings.Contains(text, substring2)
	fmt.Println(result_contains)

	// replace
	text_1 := "car"
	fmt.Println("Old String:", text_1)

	// replace r with t
	replacedText := strings.Replace(text_1, "r", "t", 1)
	fmt.Println("New String:", replacedText)

	// replace 2 r with 2 a
	out_string := strings.Replace("Programiz", "r", "R", 2)
	fmt.Println("Replaced String:", out_string)

	text1 := "go is fun to learn"
	// convert to uppercase
	text1 = strings.ToUpper(text1)
	fmt.Println(text1)

	text2 := "I LOVE GOLANG"
	// convert to lowercase
	text2 = strings.ToLower(text2)
	fmt.Println(text2)

	// split string to get a slice as output
	var msg = "I Love Golang"

	// split string from space " "
	splittedString := strings.Split(msg, " ")
	fmt.Println(splittedString)

	// repeat
	fmt.Println(strings.Repeat("Go", 3))

	// In Go, we can also represent strings using the tick mark notation
	//  represent string with `  `
	message := `I love Go Programming`

	fmt.Println(message)

}

/*
 Derived types:
	(a) Pointer
	(b) Array
	(c) Structure
	(d) Union
	(e) Function
	f) Slice
	g) Interface
	h) Map
	i) Channel

 Array types and structure types are collectively referred to as aggregate types.
*/
