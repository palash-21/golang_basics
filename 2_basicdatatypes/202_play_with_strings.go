package main

import (
	"fmt"
	"strings"
)

func main() {

	// String from a slice using join
	// create a string slice
	words := []string{"I", "love", "Golang"}

	// join each element of the slice
	message := strings.Join(words, " ")
	fmt.Println(message)

	fmt.Println(strings.Repeat("-", 10))
	// loop through string using range keyword
	text := "Golang"
	// for range loop to iterate through a string
	for _, character := range text {
		fmt.Printf("%c\n", character)
	}

	fmt.Println(strings.Repeat("-", 10))
	// Problem : get char in a string using index
	// initializing a string
	var str string = "Prevention is better than cure"
	fmt.Println("The given string is:", str)

	// need to fetch char at an given index
	var index int = 24

	// solution 1: Using split
	chars := strings.Split(str, "")
	var result1 string = chars[index]
	fmt.Println("The character at index", index, "is:", result1)

	fmt.Println(strings.Repeat("-", 10))
	// solution 2: Using rune
	// We will use the function getChar
	var result2 string = string(getChar(str, index))
	fmt.Println("The character at index", index, "is:", result2)

	fmt.Println(strings.Repeat("-", 10))
	/* rune in Go is a data type that stores codes that represent Unicode characters.
	Unicode is actually the collection of all possible characters present in the whole world.
	In Unicode, each of these characters is assigned a unique number called the Unicode code point.
	This code point is what we store in a rune data type.

	Each rune actually refers to only one Unicode code point, meaning one rune represents a single character.
	Furthermore, Go does not have a char data type, so all variables initialized with a character
	would automatically be typecasted into int32, which, in this case, represents a rune. */

	// printing rune values for all chars of a string
	var my_str string = "ABCD"
	r_array := []rune(my_str)
	fmt.Printf("Array of rune values for A, B, C and D: %U\n", r_array)

}

// function to get characters from strings using index
func getChar(str string, index int) rune {
	return []rune(str)[index]
}
