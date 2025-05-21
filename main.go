package main

import "fmt"

// type Person struct {
// 	name string
// 	age int
// 	hight int
// 	class string
// }

// func (p Person) receiver() {
//   fmt.Println("Name :", p.name)
//   fmt.Println("Age :", p.age)
//   fmt.Println("Hight :", p.hight)
//   fmt.Println("Class :", p.class)
// }


func main() {
ab := make([]int, 8)

fmt.Println(ab)
fmt.Println(len(ab))
fmt.Println(cap(ab))

}

func init() {
	fmt.Println("Hello I am first run")
}
  