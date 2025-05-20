package main

import "fmt"

type User struct{
	name string    //member variable  or property
	age int
}

func printUsrDetails(usr User) {
	fmt.Println("Name :", usr.name)  
	fmt.Println("Age :", usr.age) 
}

// এটা একটা Receiver Function
func (us User) printDetails() {
	fmt.Println("Name :", us.name)  
	fmt.Println("Age :", us.age) 
}

func (usr User) call(a int){
	fmt.Println(usr.name)
	fmt.Println(a)
}

func main() {
	var user1 User
 
	user1 = User{ //instaniate
	name: "Rahat",
	age: 22,
 }

//  printUsrDetails(user1)
user1.printDetails()

user1.call(12)

 var user2 User

 user2 = User{ // Instance or object
	name: "sara",
	age: 22,
 }

 printUsrDetails(user2)

}
