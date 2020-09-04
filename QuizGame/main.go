package main

import (
	"os"
)

func main() {
	args := os.Args[1:]
	filename := "problems.csv"

	// fmt.Println(len(args))
	// fmt.Println(args[0])

	if len(args) > 0 {
		filename = args[0]
	}
	quiz(csvReader(filename))
}
