package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("./words.txt")
	_ = data

	wordCount := 0

	for _, x := range data {
		if x == ' ' || x == '\n' {
			wordCount++
		}
	}

	fmt.Println("Word count:", wordCount)

}
