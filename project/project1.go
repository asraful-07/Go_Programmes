package main

import "fmt"

func main() {
//  fmt.Println("Welcome to the quiz game")

 var name string 
 fmt.Printf("Enter Your Name: ")
 fmt.Scan(&name)
 fmt.Printf("Hello %v, Welcome to the game!\n", name)

 var age uint 
 fmt.Printf("Enter Your Age: ")
 fmt.Scan(&age)

 if age > 10 {
   fmt.Println("Yay you can play game")
 } else {
   fmt.Println("You can't play game!")
   return
 }

 score := 0
 num_que := 2

 fmt.Printf("What is better, RTX 3080 or RTX 3090? ")
 var answer1 string
 var answer2 string
 fmt.Scan(&answer1, &answer2)
 
 if answer1 + " " + answer2 == "RTX 3090" {
   fmt.Println("correct!")
   score++
 } else if answer1 + " " + answer2 == "rtx 3090" {
   fmt.Println("correct!")
   score++
 } else {
   fmt.Println("Incorrect!")
 }

 fmt.Printf("How cores docs the intel 9 rtx 3900x have? ")
 var cores uint 
 fmt.Scan(&cores)

 if cores == 12 {
  fmt.Println("correct!")
   score++
 } else {
  fmt.Println("Incorrect!")
 }

 fmt.Printf("Your Scored %v out of %v.\n", score, num_que)
 present := (float64(score) / float64(num_que)) * 100
 fmt.Printf("Your Scored: %v%%.", present)

}

func init() {
   fmt.Println("Welcome to the quiz game")
}
