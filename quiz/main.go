package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type Problem struct {
	question string
	answer   string
}

func main() {

	timeLimit := flag.Int("limit", 10, "the time limit for quiz in seconds")
	flag.Parse()

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
			answer:   strings.TrimSpace(row[1]),
		}

		problems = append(problems, pro)
	}

	correct := 0
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	answerCh := make(chan string)

problemLoop:
	for i := 0; i < len(problems); i++ {
		fmt.Printf("give an answer of %s : ", problems[i].question)

		go func() {
			var answer string
			fmt.Scan(&answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println("\nTime's Up")
			break problemLoop
		case answer := <-answerCh:
			if problems[i].answer == answer {
				correct++
			}
		}
	}

	fmt.Printf("You answered %d correctly out of %d problems \n", correct, len(problems))
}
