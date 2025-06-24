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
	// Get quiz file name from user
	var inputFileName string

	fmt.Print("What would you like to name the new quiz file? ")
	fmt.Scanln(&inputFileName)

	var quizFileName string
	
	// Handle if either the user did not add the ".csv" extension themselves or simply making a new.csv if no characters were input
	if (strings.Contains(inputFileName, ".csv")) {
		quizFileName = fmt.Sprintf("%v", inputFileName)

	} else if (strings.Trim(inputFileName, " ") == "") {
		quizFileName = "new.csv"

	} else {
		quizFileName = fmt.Sprintf("%v.csv", inputFileName)

	}

	// Create the file
	newQuizFile, quizFileErr := os.Create(quizFileName)
	if (quizFileErr != nil) {
		log.Fatal(quizFileErr)
	}

	// Initialize a csv writer
	quizWriter := csv.NewWriter(newQuizFile)

	// The condition that allows for adding questions
	questionAdding := true

	var question string
	var answer string

	for (questionAdding) {
		// Get the question and its associated answer from user input
		var addQuestion string
		var addAnother string
		fmt.Print("Question: ")
		fmt.Scanln(&question)

		fmt.Print("Answer: ")
		fmt.Scanln(&answer)
		fmt.Println()

		fmt.Print("Confirm question? (y/n) ")
		fmt.Scanln(&addQuestion)
		fmt.Println()

		// Confirms if the user correctly typed the question and answer, to add it to the writer
		if (ynConfirm(addQuestion)) {
			questionSet := []string{question, answer}
			quizWriter.Write(questionSet)
		}

		// Asks if the user would like to continue the loop and add another question.
		fmt.Printf("Would you like to add another question? (y/n) ")
		fmt.Scanln(&addAnother)
		fmt.Println()

		// Sets the bool for the loop condition to false when the user is done
		if (!ynConfirm(addAnother)) {
			questionAdding = false
		}
	} 
	// Post-loop, flush the records to the csv and close it
	quizWriter.Flush()
	fmt.Printf("Questions successfully written to %v\n", newQuizFile.Name())

	defer newQuizFile.Close()
}