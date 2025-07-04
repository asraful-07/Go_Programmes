package main

import (
	"fmt"
	"net/http"
)

func handleUsers(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"Hello bro how are you.")
	fmt.Fprintln(w,"Hello bro how are you. And you bro?")
}

func handleProduct(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"This is product.")
	fmt.Fprintln(w,"what is product?")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/users", handleUsers)
	mux.HandleFunc("/product", handleProduct)

	fmt.Println("User Server is running 8080")

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		fmt.Println("Start Server Problems")
	}

}