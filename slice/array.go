package main

import "fmt"

func main() {
	arr := [10]string{"a","s","d","f","g","h","j","k","l","z"}
	fmt.Println(arr)
	

	a := arr[1:6]
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
   
	a1 := a[1:4]
	fmt.Println(a1)
	fmt.Println(len(a1))
	fmt.Println(cap(a1))
}
