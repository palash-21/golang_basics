/* Maps are used to store data values in key:value pairs.
A map is an unordered and changeable collection that does not allow duplicates.
The length of a map is the number of its elements. You can find it using the len() function.
The default value of a map is nil.
Maps hold references to an underlying hash table */

package main

import (
	"fmt"
	"strings"
)

func main() {

	// create map with var keyword
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}

	// create map with :=
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	// The order of the map elements defined in the code is different from the way that they are stored.
	// The data are stored in a way to have efficient data retrieval from the map
	// So when we print the map, the order is different
	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)

	fmt.Println(strings.Repeat("-", 10))
	// create map with make keyword
	var c = make(map[string]string) // The map c is empty now
	c["brand"] = "Ford"
	c["model"] = "Mustang"
	c["year"] = "1964" // a is no longer empty

	fmt.Println(strings.Repeat("-", 10))
	d := make(map[string]int)
	d["Oslo"] = 1
	d["Bergen"] = 2
	d["Trondheim"] = 3
	d["Stavanger"] = 4

	fmt.Printf("a\t%v\n", c)
	fmt.Printf("b\t%v\n", d)

	fmt.Println(strings.Repeat("-", 10))
	// The make()function is the right way to create an empty map.
	// If you make an empty map in a different way and write to it, it will causes a runtime panic
	// create an empty map without using make
	var e = make(map[string]string)
	var f map[string]string

	fmt.Println(e == nil)
	fmt.Println(f == nil) // notice that f is actually nil

	fmt.Println(strings.Repeat("-", 10))
	// Accessing elements using key
	fmt.Println(a["year"])

	fmt.Println(strings.Repeat("-", 10))
	fmt.Println(a)
	a["year"] = "1970" // Updating an element
	a["color"] = "red" // Adding an element
	fmt.Println(a)

	fmt.Println(strings.Repeat("-", 10))
	// Deleting elements using key
	delete(a, "color")
	fmt.Println(a)

	fmt.Println(strings.Repeat("-", 10))
	// You can check if a certain key exists in a map and check its value
	var g = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day": ""}

	val1, ok1 := g["brand"] // Checking for existing key and its value
	val2, ok2 := g["color"] // Checking for non-existing key and its value
	val3, ok3 := g["day"]   // Checking for existing key and its value
	_, ok4 := g["model"]    // Only checking for existing key and not its value

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
	fmt.Println(ok4)
	/* The ok2 variable is used to find out if the key exist or not.
	Because we would have got the same value if the value of the "color" key was empty.
	This was the case for val3. */

	fmt.Println(strings.Repeat("-", 10))
	/* Maps are references to hash tables.
	If two map variables refer to the same hash table, changing the content of one variable affect the content of the other */
	var x = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	y := x

	fmt.Println(x)
	fmt.Println(y)

	y["year"] = "1970"
	fmt.Println("After change to y:")
	fmt.Println(x)
	fmt.Println(y)

	fmt.Println(strings.Repeat("-", 10))
	// Using range keyword to iterate over map
	i := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	// fetching keys and values
	for k, v := range i {
		fmt.Printf("%v : %v, ", k, v)
	}

	fmt.Println("")
	fmt.Println(strings.Repeat("-", 10))
	// fetching only keys
	for k := range i {
		fmt.Printf("%v ", k)
	}

}

/* The map key can be of any data type from below :
Booleans
Numbers
Strings
Arrays
Pointers
Structs
Interfaces (as long as the dynamic type supports equality) */

/* Invalid key types are:
Slices
Maps
Functions */

/*The map values can be any type.*/
