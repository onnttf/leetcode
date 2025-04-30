package main

import (
	"fmt"
	"os"

	"leetcode/script"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <slug>")
		fmt.Println("Example: go run main.go two-sum")
		return
	}

	slug := os.Args[1]
	script.Generate(slug)
}
