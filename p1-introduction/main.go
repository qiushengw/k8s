package main

import (
"fmt"
"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Today is Sunny!\n")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Simple weather server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
