package main

import (
	"fmt"
	"leetcode/script"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <slug>")
		fmt.Println("Example: go run main.go two-sum")
		return
	}

	slug := os.Args[1]
	script.GenerateCode(slug)
	script.GenerateReadme()
}
