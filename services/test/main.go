package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})
	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":80", nil); err != nil {
		fmt.Println(err)
	}
}
