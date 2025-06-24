package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func ynConfirm(s string) bool {
	if (strings.ToLower(s) == "y") {
		return true
	} else {
		return false
	}
}

func main() {
	var inputFileName string

	fmt.Print("What would you like to name the new quiz file? ")
	fmt.Scanln(&inputFileName)

	var quizFileName string
	
	if (strings.Contains(inputFileName, ".csv")) {
		quizFileName = fmt.Sprintf("%v", inputFileName)
	} else {
		quizFileName = fmt.Sprintf("%v.csv", inputFileName)
	}

	newQuizFile, quizFileErr := os.Create(quizFileName)
	if (quizFileErr != nil) {
		log.Fatal(quizFileErr)
	}
	
	quizWriter := csv.NewWriter(newQuizFile)

	questionAdding := true

	var question string
	var answer string

	for (questionAdding) {
		var addQuestion string
		var addAnother string
		fmt.Print("Question: ")
		fmt.Scanln(&question)

		fmt.Print("Answer: ")
		fmt.Scanln(&answer)
		fmt.Println()

		fmt.Printf("Q: %v\nA: %v\n\nConfirm question? (y/n) ", question, answer)
		fmt.Scanln(&addQuestion)
		fmt.Println()

		if (ynConfirm(addQuestion)) {
			questionSet := []string{question, answer}
			quizWriter.Write(questionSet)
		}

		fmt.Printf("Would you like to add another question? (y/n) ")
		fmt.Scanln(&addAnother)
		fmt.Println()

		if (!ynConfirm(addAnother)) {
			questionAdding = false
		}

		quizWriter.Flush()

		fmt.Printf("Questions successfully written to %v\n", newQuizFile.Name())
	} 
}