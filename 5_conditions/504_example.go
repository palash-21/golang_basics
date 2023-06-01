/* In this exercise you'll be writing code to analyze the production in a car factory.
The cars are produced on an assembly line.
The assembly line has a certain speed, that can be changed.
The faster the assembly line speed is, the more cars are produced.
However, changing the speed of the assembly line also changes the number of cars that are produced successfully, that is cars without any errors in their production.
*/

/* 1. Implement a function that takes in the number of cars produced per hour and the success rate and calculates the number of successful cars made per hour.
The success rate is given as a percentage, from 0 to 100 */

/* 2. Implement a function that takes in the number of cars produced per hour and the success rate and calculates how many cars are successfully produced each minute */

/* 3. Each car normally costs $10,000 to produce individually, regardless of whether it is successful or not. But with a bit of planning, 10 cars can be produced together for $95,000.
For example, 37 cars can be produced in the following way: 37 = 3 x groups of ten + 7 individual cars
So the cost for 37 cars is: 3*95,000+7*10,000=355,000
Implement the function CalculateCost that calculates the cost of producing a number of cars, regardless of whether they are successful */

// Solution:

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.

package main

import "fmt"

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	//panic("CalculateWorkingCarsPerHour not implemented")
	result := float64(productionRate) * (successRate / 100)
	// fmt.Println(result)
	return result
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	//panic("CalculateWorkingCarsPerMinute not implemented")

	perhour := CalculateWorkingCarsPerHour(productionRate, successRate)
	perMin := int(perhour / 60)
	return perMin
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	//panic("CalculateCost not implemented")
	var cars_count1 int
	if carsCount > 10 {
		cars_count1 = carsCount / 10
	} else if carsCount == 10 {
		cars_count1 = 1
	} else {
		cars_count1 = 0
	}

	var cost uint = uint(cars_count1*95000 + 10000*(carsCount-cars_count1*10))
	return cost
}

func main() {

	// testing func 1
	fmt.Println(CalculateWorkingCarsPerHour(1200, 90))

	// testing func 2
	fmt.Println(CalculateWorkingCarsPerMinute(1200, 90))

	// testing func 3
	fmt.Println(CalculateCost(100))
	fmt.Println(CalculateCost(10))
	fmt.Println(CalculateCost(5))
}
