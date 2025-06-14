package main

import "fmt"

type Cars struct {
	name string  
	price float64
	color string
	body string
	res [4]int
	sum []string
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
	res: [4]int{1,2,3,4},
	sum: []string{"hi", "how", "are", "you"},
}

fmt.Println(person)
modifyObject(&person)
fmt.Println(person)

arr := [4]string{"A", "B", "C", "D"}

fmt.Printf("string%s\n", arr)
modifyArray(&arr)
fmt.Println(arr)

}
