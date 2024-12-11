package main

import (
	"encoding/csv"
	"os"
)

type Problem struct {
	question string
	answer   string
}

func main() {

	var problems []Problem

	file, err := os.Open("problems.csv")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	reader.FieldsPerRecord = -1

	data, err := reader.ReadAll()

	if err != nil {
		panic(err)
	}

	for _, row := range data {

		pro := Problem{
			question: row[0],
			answer:   row[1],
		}

		problems = append(problems, pro)
	}
}
