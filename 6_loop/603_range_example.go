/* The range keyword is used in for loop to iterate over items of an array, slice, channel or map. With array and slices, it returns the index of the item as integer. With maps, it returns the key of the next key-value pair. Range either returns one value or two. If only one value is used on the left of a range expression, it is the 1st value in the following table.

Range expression	    |     1st Value      |    	2nd Value(Optional)
Array or slice a [n]E	|    index i int	 |         a[i] E
String s string type	|    index i int	 |        rune int
map m map[K]V	        |      key k K	     |       value m[k] V
channel c chan E	    |    element e E	 |          none
*/

package main

import (
	"fmt"
	"strings"
)

func main() {

	// string
	var s = "Golang"

	// iterate over string
	for ind, val := range s {
		fmt.Println("String item", ind, "is", val)
	}

	fmt.Println(strings.Repeat("-", 10))
	/* create a slice */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	/* print the numbers */
	for i := range numbers {
		fmt.Println("Slice item", i, "is", numbers[i])
	}

	fmt.Println(strings.Repeat("-", 10))
	/* create a map*/
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}

	/* print map using keys*/
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	/* print map using key-value*/
	for country, capital := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", capital)
	}
}
