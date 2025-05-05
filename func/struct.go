package main

import "fmt"

// var (
//   a = 10 
//   b = 40
// )

type User struct{
	name string    //member variable  or property
	age int
}

func 

func main() {
	var user1 User
 
	user1 = User{ //instaniate
	name: "Rahat",
	age: 22,
 }

 fmt.Println("Name :", user1.name)  
 fmt.Println("Age :", user1.age) 

 var user2 User

 user2 = User{ // Instance or object
	name: "sara",
	age: 22,
 }

 
 fmt.Println("Name :", user2.name)
 fmt.Println("Age :", user2.age)
}
