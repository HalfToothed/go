package cmd

import (
	"fmt"
	"strconv"
)

func Add(first string, second string) (result string) {
	num1, err := strconv.ParseFloat(first, 64)
	if err != nil {
		fmt.Println("Error: First value is invalid")
		return
	}
	num2, err := strconv.ParseFloat(second, 64)
	if err != nil {
		fmt.Println("Error: Second value is invalid")
		return
	}
	return fmt.Sprintf("%f", num1+num2)
}

func Subtract(from string, subtract string) (result string) {
	num1, err := strconv.ParseFloat(from, 64)
	if err != nil {
		fmt.Println("Error: First value is invalid")
		return
	}
	num2, err := strconv.ParseFloat(subtract, 64)
	if err != nil {
		fmt.Println("Error: Second value is invalid")
		return
	}
	return fmt.Sprintf("%f", num1-num2)
}
