package main

import (
	"fmt"
	"net/http"
)

func main() {
	response, err := http.Get("https://google.com")
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	fmt.Printf("status code: %d", response.StatusCode)
}
