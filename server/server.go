package main

import (
	"fmt"
	"net/http"
)

func handleUsers(u http.ResponseWriter,t *http.Request){
	fmt.Fprintln(u,"Hello bro how are you.")
	fmt.Fprintln(u,"Hello bro how are you. And you bro?")
}

func handleProduct(u http.ResponseWriter,t *http.Request){
	fmt.Fprintln(u,"This is product.")
	fmt.Fprintln(u,"what is product?")
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