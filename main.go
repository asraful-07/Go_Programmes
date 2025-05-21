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
