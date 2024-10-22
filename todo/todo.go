package main

import (
	"bufio"
	"fmt"
	"os"
)

var data []string

func main() {

	var input int

	fmt.Print("A Simple TO-DO List app \n")

	for {

		fmt.Print("\nChoose an option (1-4) \n")
		Options()

		fmt.Scanln(&input)

		if input == 1 {
			Add()

		} else if input == 2 {
			List()

		} else if input == 3 {
			Delete()
		} else if input == 4 {
			fmt.Print("Exiting the application. Goodbye!")
			break
		}
	}
}

func Options() {
	fmt.Print("1. Add a new task \n")
	fmt.Print("2. View all the tasks \n")
	fmt.Print("3. Delete a task \n")
	fmt.Print("4. Exit \n\n")
}

func Add() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Write the Task: \n")

	addInput, _ := reader.ReadString('\n')

	data = append(data, addInput)

	fmt.Print("Task added sucessfully :) \n")
}

func List() {

	fmt.Print("The task list: \n")

	for index, value := range data {
		fmt.Printf("%d. %s", index+1, value)
	}
}

func Delete() {

	var index int

	fmt.Print("write the index of the task to be deleted \n")

	fmt.Scanln(&index)

	index -= 1

	if index >= 0 && index < len(data) {
		data = append(data[:index], data[index+1:]...)
		fmt.Print("task deleted \n")
	} else {
		fmt.Print("Index out of range")
	}

}
