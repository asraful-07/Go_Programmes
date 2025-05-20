package main

import "fmt"

type Person struct {
	name string
	age int
	hight int
	class string
}

func (p Person) receiver() {
  fmt.Println("Name :", p.name)
  fmt.Println("Age :", p.age)
  fmt.Println("Hight :", p.hight)
  fmt.Println("Class :", p.class)
}

func main() {
  Person1 := Person{
	name : "Rahat",
	age : 22,
	hight : 167,
	class : "CSE",
  }

  fmt.Println("Name :", Person1.name)
  fmt.Println("Age :", Person1.age)
  fmt.Println("Hight :", Person1.hight)
  fmt.Println("Class :", Person1.class)

  var Person2 Person 

  Person2 = Person{
	name : "Sara",
	age : 22,
	hight : 160,
	class : "CSE",
  }
Person2.receiver()

}

func init() {
	fmt.Println("Hello I am first run")
}
  