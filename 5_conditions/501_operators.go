package main

import (
	"fmt"
)

func main() {
	// Arithmetic operators
	var x, y = 50, 5
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y) // For integer division, the remainder is dropped (e.g. 5 / 2 == 2).
	fmt.Println(x % y)
	y++
	x--
	fmt.Println(x, y)

	// Assignment
	var a, b = 56, 33
	fmt.Println(a)
	a += 2
	fmt.Println(a)
	a -= 2
	fmt.Println(a)
	a *= 2
	fmt.Println(a)
	a /= 2
	fmt.Println(a)
	b %= 2
	fmt.Println(b)
	b &= 2
	fmt.Println(b)
	b |= 2
	fmt.Println(b)
	b ^= 2
	fmt.Println(b)
	b = 10
	b <<= 2
	fmt.Println(b)
	b >>= 2
	fmt.Println(b)

	// Comparison
	fmt.Println(a == b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
	fmt.Println(a != b)

	// Logical
	fmt.Println(a < b && x > y)
	fmt.Println(a < b || x > y)
	fmt.Println(!(a < b && x > y))

	// Bitwise
	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	fmt.Println(a << 2)
	fmt.Println(a >> 2)

}
