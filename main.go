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

 fmt.Printf("What is better, RTX 3080 or RTX 3090? ")
 var answer1 string
 var answer2 string
 fmt.Scan(&answer1, &answer2)
 
 if answer1 + " " + answer2 == "RTX 3090" {
   fmt.Println("correct")
 } else {
   fmt.Println("Incorrect!")
 }

}

func init() {
   fmt.Println("Welcome to the quiz game")
}
