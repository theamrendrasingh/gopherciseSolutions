package main

import (
	"flag"
)

var (
	file  = flag.String("csv", "problems.csv", "a csv file in the format 'question,answer'")
	limit = flag.Int("limit", 10, "the time limit for the quiz in seconds")
)

func main() {
	filename := *file
	quizTime := *limit

	quiz(csvReader(filename), quizTime)
}
