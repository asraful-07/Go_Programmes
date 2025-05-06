package main

import "fmt"

type Off struct{
	Name string
	Age  int
}

var add = [4]int{4, 5, 6, 7}

func modifyArray(num *[10]int) {
    (*num)[3] = 90
}

func modifyObject(p *Off) {
	p.Name = "GG"
}

func main() {

	ff := Off{
		Name: "FF",
		Age: 10,
	}
	fmt.Println(ff)

	modifyObject(&ff)
	fmt.Println(ff)

    arr := [10]int{7, 8, 6, 6, 6}
    fmt.Println("Before:", arr)
    fmt.Println("Add:", add)

    modifyArray(&arr)

    fmt.Println("After:", arr)
}
