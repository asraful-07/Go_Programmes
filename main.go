package main

import "fmt"

var a = 10

func main() {
    age := 30

    if age > 18 {
        a := 47        //  শুধু এই if ব্লকের ভিতরে গ্লোবাল `a` কে shadow করে
        fmt.Println(a) // ➜ 47
    }

    fmt.Println(a)     // ➜ 10 (গ্লোবাল `a` প্রিন্ট করে)
}
