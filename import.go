package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, Go Standard Library!")

	resp, err := http.Get("https://www.example.com")
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Response Status:", resp.Status)
}
