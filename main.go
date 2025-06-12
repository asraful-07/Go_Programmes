package main

import "fmt"

type Cars struct {
	name string 
	price float64
	color string
	body string
}

func (p Cars) Receiver() {
	fmt.Println(p.name)
	fmt.Println(p.price)
}

func main() {
person := Cars{
	name: "tesla",
	price: 46.276,
	color: "blue",
	body: "tata",
}

person.Receiver()


}
