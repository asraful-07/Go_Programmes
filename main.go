package main

import "fmt"

type Cars struct {
	name string  
	price float64
	color string
	body string
}

func modifyObject(p *Cars) {
  p.color = "red"
}

func main() {
person := Cars{
	name: "tesla",
	price: 46.276,
	color: "blue",
	body: "tata",
}
fmt.Println(person)
modifyObject(&person)
fmt.Println(person)

}
