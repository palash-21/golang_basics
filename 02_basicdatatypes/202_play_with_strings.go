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

}

// function to get characters from strings using index
func getChar(str string, index int) rune {
	return []rune(str)[index]
}
