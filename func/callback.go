package main

import "fmt"

// এটা হলো callback function
func sayHello(name string) {
    fmt.Println("Hello", name)
}

// এটা হলো যেটা callback function কে আর্গুমেন্ট হিসেবে নেয়
func greet(callback func(string)) {
    callback("Rahat")
}

func main() {
    // এখানে sayHello ফাংশনকে callback হিসেবে পাঠানো হচ্ছে
    greet(sayHello)
}
