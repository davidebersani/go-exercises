package main

import "fmt"

func main() {
	// Some comment
	fmt.Println("Starting Textio server")

	var username string = "davide"
	var password string = "432521"

	fmt.Println("Authorization: Basic", username+":"+password)
}
