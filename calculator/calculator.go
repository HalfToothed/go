package main

import "fmt"

func main() {

	var operation int

	var input_1 int
	var input_2 int

	fmt.Print("Welcome to my Calculator \n")
	fmt.Print("Choose Operation 1-Add, 2-Subtraction, 3-Multiplication, 4-Division \n")

	fmt.Scanln(&operation)

	fmt.Scan(&input_1, &input_2)

	switch operation {
	case 1:
		fmt.Print("Answer: ", add(input_1, input_2))
	case 2:
		fmt.Print("Answer: ", sub(input_1, input_2))
	case 3:
		fmt.Print("Answer: ", multiply(input_1, input_2))
	case 4:
		fmt.Print("Answer: ", divide(input_1, input_2))
	default:
		fmt.Print("not a valid operation")
	}
}

func add(x int, y int) int {
	sum := x + y
	return sum
}

func sub(x int, y int) int {
	return x - y
}

func multiply(x int, y int) int {
	return x * y
}

func divide(x int, y int) int {
	return x / y
}
