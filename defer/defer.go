package main

import "fmt"

// ✅ এখানে result একটা named return variable
func add() (result int) {
	fmt.Println("first", result) // শুরুতে result এর মান 0, তাই এটা প্রিন্ট হবে

	// এটা একটা ফাংশন যেটা পরে (defer) চালানো হবে
	show := func() {
		result = result + 10 // result এর মান ১০ যোগ করে আপডেট হচ্ছে
		fmt.Println("defer", result) // defer block থেকে প্রিন্ট
	}

	defer show() // 🔁 এই ফাংশনটা ফাংশন শেষ হওয়ার সময় চলবে

	result = 5 // এখন result = 5 সেট করা হলো
	fmt.Println("second", result) // প্রিন্ট হবে second 5

	return result // return করার ঠিক আগে defer কাজ করবে এবং result আপডেট হবে
}

// ❌ এখানে result হচ্ছে সাধারণ local variable, named return না
func calc() int {
	result := 0 // শুরুতে result = 0
	fmt.Println("first", result)

	show := func() {
		result = result + 10 // এটা local variable modify করে
		fmt.Println("defer", result) // defer block থেকে প্রিন্ট
	}

	defer show() // 🔁 ফাংশন শেষ হওয়ার সময় চলবে

	result = 5
	fmt.Println("second", result)

	return result // এখানে return value আগে ঠিক হয়ে যায়, তারপর defer চলে
}

func main() {
	a := add()
	fmt.Println("main first", a) // এখানে প্রিন্ট হবে 15 কারণ defer result কে modify করে

	b := calc()
	fmt.Println("main second", b) // এখানে প্রিন্ট হবে 5 কারণ return value আগে fix হয়
}
