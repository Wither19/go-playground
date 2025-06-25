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

	var confirmOverwrite string
	quizPath, quizPathErr := os.Stat(quizFileName)
	if (quizPathErr != nil) {

	} else {
		fmt.Printf("%v already exists in this directory. Would you like to overwrite it? (y/n) ", quizPath.Name())
		fmt.Scanln(&confirmOverwrite)

		if (!ynConfirm(confirmOverwrite)) {
			fmt.Println("Exiting...")
			os.Exit(0)
		}
	}

	// Create the file
	newQuizFile, quizFileErr := os.Create(quizFileName)
	if (quizFileErr != nil) {
		log.Fatal(quizFileErr)
	} else if (ynConfirm(confirmOverwrite)) {
		fmt.Printf("%v successfully overwritten\n", quizFileName)
	} else {
		fmt.Printf("%v successfully created\n", quizFileName)
	}

	// Initialize a csv writer
	quizWriter := csv.NewWriter(newQuizFile)

	// The condition that allows for adding questions
	questionAdding := true

	recordCount := 0

	var question string
	var answer string

	for (questionAdding) {
		// Get the question and its associated answer from user input
		var addQuestion string
		var addAnother string
		fmt.Printf("Question %d: ", recordCount + 1)
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

			recordCount += 1
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
	qString := "questions"

	if (recordCount == 1) {
		qString = "question"
	}

	fmt.Printf("%d %v successfully written to %v\n", recordCount, qString, quizFileName)

	defer newQuizFile.Close()
}