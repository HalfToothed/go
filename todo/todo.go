package main

import "fmt"

func main() {

	var input int

	fmt.Print("A Simple TO-DO List app \n")
	fmt.Print("1. Add a new task \n")
	fmt.Print("2. View all the tasks \n")
	fmt.Print("3. Delete a task \n")
	fmt.Print("4. Exit")

	for {

		fmt.Print("\n\nChoose an option (1-4) \n")

		fmt.Scanln(&input)

		if input == 1 {
			fmt.Print("Add")

		} else if input == 2 {
			fmt.Print("list")

		} else if input == 3 {
			fmt.Print("delete")
		} else if input == 4 {
			fmt.Print("Exiting the application. Goodbye!")
			break
		}
	}

}
