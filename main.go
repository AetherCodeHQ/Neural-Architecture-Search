package main

import (
	"fmt"
	"os"
)

// neural_architecture_search - Auto neural network optimization
func neural_architecture_search(path string) {
	fmt.Println("========================================")
	fmt.Println("  Neural-Architecture-Search")
	fmt.Println("  Auto neural network optimization")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	neural_architecture_search(path)
}
