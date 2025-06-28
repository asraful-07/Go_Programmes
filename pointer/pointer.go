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

// &x হলো x এর ঠিকানা (address) রাখে
// *p দিয়ে আমরা x এর ঠিকানায় গিয়ে তার আসল মান পড়ি বা সেট করি।

package main

import "fmt"

func main() {
    var x int = 10
    var p *int = &x  // x এর address নিচ্ছি

    fmt.Println("x এর মান:", x)       // 10
    fmt.Println("x এর address:", &x)  // x এর মেমোরি address
    fmt.Println("p এর মান:", p)       // p তে x এর address আছে
    fmt.Println("*p:", *p)            // *p মানে, সেই address এ গিয়ে মান পড়া, মানে 10

    *p = 20  // পয়েন্টার ব্যবহার করে x এর মান পরিবর্তন

    fmt.Println("x এর নতুন মান:", x)  // এখন x এর মান ২০ হয়ে গেছে
}

