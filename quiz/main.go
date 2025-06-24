package main

// https://github.com/gophercises/quiz

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	// Initialize a flag to change the file opened for the questions
	problemFile := flag.String("file", "problems.csv", "The CSV for the set of problems used")
	flag.Parse()

	quizFile, fileOpenErr := os.Open(*problemFile)

	if (fileOpenErr != nil) {
		log.Fatal(fileOpenErr)
	}

	reader := csv.NewReader(quizFile)

	problems, problemsErr := reader.ReadAll()

	if (problemsErr != nil) {
		log.Fatal(problemsErr)
	} 

	correctAnswerCount := 0

		for i, problem := range problems {
			correctAnswerCount += presentProblem(i + 1, problem)

			if (i == len(problems) - 1) {
				fmt.Printf("\nYour results:\n%d/%d answered correctly\n", correctAnswerCount, len(problems))
			}
		}
	defer quizFile.Close()
}

func presentProblem(questionNum int, problemSet []string) int {
	isCorrect := 0

	problem, answer := problemSet[0], problemSet[1]
	
	var userAnswer string

	fmt.Printf("#%d. %v ", questionNum, problem)
	fmt.Scanln(&userAnswer)

	userAnswer = strings.Trim(strings.ToLower(userAnswer), " ")

	if (userAnswer == answer) {
		fmt.Println("\nCorrect!")
		isCorrect = 1

	} else {
		fmt.Println("\nIncorrect")
	}

	return isCorrect

}