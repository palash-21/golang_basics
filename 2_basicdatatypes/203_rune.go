/* A Go string is a read-only slice of bytes.

Rune in Go is a data type that stores codes that represent Unicode characters.
Unicode is actually the collection of all possible characters present in the whole world.
In Unicode, each of these characters is assigned a unique number called the Unicode code point.
This code point is what we store in a rune data type.

Each rune actually refers to only one Unicode code point, meaning one rune represents a single character.
Furthermore, Go does not have a char data type, so all variables initialized with a character
would automatically be typecasted into int32, which, in this case, represents a rune.

Since strings are equivalent to []byte, this will produce the length of the raw bytes stored within.
Indexing into a string produces the raw byte values at each index.
This loop generates the hex values of all the bytes that constitute the code points in a string.
*/

package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	// printing rune values for all chars of a string
	var my_str string = "ABCD"
	r_array := []rune(my_str)
	fmt.Printf("Array of rune values for A, B, C and D: %U\n", r_array)

	fmt.Println(strings.Repeat("-", 10))
	// To count how many runes are in a string, we can use the utf8 package

	const s = "Golang"

	fmt.Println("Len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i]) // note that %x will give hex dump of byte values
	}
	fmt.Println(strings.Repeat("-", 10))

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}
	// Values enclosed in single quotes are rune literals.
	// We can compare a rune value to a rune literal directly.

	fmt.Println(strings.Repeat("-", 10))
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}

}

func examineRune(r rune) {

	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
