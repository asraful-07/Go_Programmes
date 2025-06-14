package main

import "fmt"

type Cars struct {
	name string  
	price float64
	color string
	body string
	rul []int
}

func modifyObject(p *Cars) {
  p.color = "red"
}

func modifyArray(a *[4]string) {
 (*a)[2] = "H"
}

func main() {
person := Cars{
	name: "tesla",
	price: 46.276,
	color: "blue",
	body: "tata",
	rul: []int{1,2,3,4,5,6},
}

fmt.Println(person)
modifyObject(&person)
fmt.Println(person.rul[3])

arr := [4]string{"A", "B", "C", "D"}

fmt.Printf("string%s\n", arr)
modifyArray(&arr)
fmt.Println(arr)

}
