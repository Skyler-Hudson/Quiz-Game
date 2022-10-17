package main

import "fmt"

func greeting() {
	fmt.Print("Please enter your name: ")
	var name string
	fmt.Scan(&name)

	fmt.Println("Hello", name)
}

func quiz() {
	var score int = 0
	var totalQuestions int = 2

	// Question 1
	fmt.Println("Question 1: What is the color of the sky? ")
	var answer string
	fmt.Scan(&answer)
	if answer == "blue" {
		score++
		fmt.Println("Correct")
	} else if answer == "Blue" {
		fmt.Println("Correct")
		score++
	} else {
		fmt.Println("Incorrect")
	}

	// Question 2
	fmt.Println("Question 2: Is the grass green?")
	var answer2 string
	fmt.Scan(&answer2)
	if answer2 == "yes" {
		fmt.Println("Corrent!")
		score++
	} else if answer2 == "Yes" {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect :(")
	}

	// Total Score and Percentage
	fmt.Printf("Your total score was: %v out of %v questions\n", score, totalQuestions)
	var percentage float64 = (float64(score) / float64(totalQuestions)) * 100
	fmt.Printf("Your percentage was %v%%", percentage)
}

func main() {
	fmt.Println("Welcome to the Quiz Game!")
	greeting()
	quiz()
}
