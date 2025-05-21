package main

import "fmt"

func main() {
	var s []int         // এখানে s হলো একটা খালি slice, মানে এর len = 0, cap = 0
	s = append(s, 9)    // len = 1, cap = 1
	s = append(s, 7)    // len = 2, cap = 2
	s = append(s, 3)    // len = 3, cap = 4  <-- এখানে Go নিজে থেকে capacity বাড়িয়ে নেয়
	
	r := s              // এখন r এবং s একই backing array শেয়ার করে
	s = append(s, 4)    // এখন s = [1 2 3 4], len = 4, cap = 4
	s = append(s, 5)    // এখন capacity শেষ হয়ে গেছে (4 ছিল), তাই Go নতুন backing array বানায়
	
	r = append(r, 4)    // r এখনো পুরাতন backing array ব্যবহার করছে, তাই ওখানে ৪ যোগ করে
	s[0] = 10           // এটা নতুন backing array তে হচ্ছে, তাই r-এর উপর কোন প্রভাব নেই
	


    fmt.Println(s, len(s), cap(s))
    fmt.Println(r, len(r), cap(r))
}





package main

import "fmt"

func main() {

var arr []int

fmt.Println(arr)

 /*
 
 var s []int // []

s = append(s, 1, 2, 3) //[1], len: 3 , cap: 3

fmt.Println(s, len(s), cap(s)
 
 */

}
