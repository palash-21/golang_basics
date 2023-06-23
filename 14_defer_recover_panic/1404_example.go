package main

import "fmt"

func div(num1, num2 int) int {
	defer func() {
		fmt.Println(recover())
	}()
	solution := num1 / num2

	return solution
}

// func to demonstrate panic
func demPanic() {
	defer func() {
		fmt.Println(recover())
	}()
	panic("Panic")
}

func main() {
	fmt.Println(div(3, 0))
	fmt.Println(div(10, 5))
	demPanic()
}
