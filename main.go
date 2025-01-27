package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}
