package main

import "fmt"

type Person struct{
	Name string
	Age int
	Salary float32
	FavFoods [4]string
}

func modifyArray(arr *[3]int) {
    (*arr)[0] = 999
}

func modifyObject(p *Person){
	p.Age = 23
} 


func main() {
	p := Person{
		Name : "Rahat",
		Age: 22,
		Salary: 0,
		FavFoods: [4]string{"Apple", "Banana", "Tomato", "Potato"},
	}

  fmt.Println(p) //after
  modifyObject(&p)
  fmt.Println(p) //before

    nums := [3]int{10, 20, 30}
    fmt.Println("After:", nums)

    modifyArray(&nums)

    fmt.Println("Before:", nums)
}




