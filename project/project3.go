package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to quiz game!")

	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	fmt.Printf("Hello %v, Welcome to the game!\n", name)

	var gmail string
	fmt.Print("Please enter your Gmail to login: ")
	fmt.Scan(&gmail)

	if strings.HasSuffix(gmail, "@gmail.com") {
		fmt.Println("Login successful!")
		fmt.Println("Let's start our quiz game!\n")
	} else {
		fmt.Println("Please login using a valid Gmail address.")
		return
	}

	score := 0
  num_que := 6

	// Question 1
	fmt.Print("1. Who is better, Messi or Ronaldo? ")
	var answer string
	fmt.Scan(&answer)
	if strings.ToLower(answer) == "messi" {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	// Question 2
	fmt.Print("2. Who won the World Cup in 2022, Messi or Mbape? ")
	var answer1 string
	fmt.Scan(&answer1)
	if strings.ToLower(answer1) == "messi" {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	// Question 3
	fmt.Print("3. Who is the best captain of India, Dhoni or Rohit? ")
	var answer2 string
	fmt.Scan(&answer2)
	if strings.ToLower(answer2) == "dhoni" {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	// Question 4
	fmt.Println("4. Who has won IPL 5 times as captain, Dhoni or Tendulkar?")
	var answer3 string
	fmt.Scan(&answer3)
	if strings.ToLower(answer3) == "dhoni" {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	// Question 5
	fmt.Println("5. What is 6 + 8, 14 or 20?")
	var number uint8
	fmt.Scan(&number)
	if number == 14 {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	// Question 6
	fmt.Println("6. What is 12 + 8, 14 or 20?")
	var number1 uint8
	fmt.Scan(&number1)
	if number1 == 20 {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

  // Final Score
  fmt.Printf("\nThanks for playing, %v!\nYour final score is: %d out of 6\n", name, score)
  present := (float64(score) / float64(num_que)) * 100
  fmt.Printf("Your Scored: %v%%", present)
}
