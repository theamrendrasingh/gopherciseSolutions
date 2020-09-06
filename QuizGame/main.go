package main

import (
	"flag"
)

var (
	file    = flag.String("csv", "problems.csv", "a csv file in the format 'question,answer'")
	limit   = flag.Int("limit", 10, "the time limit for the quiz in seconds")
	shuffle = flag.Bool("shuffle", false, "shuffle the questions")
)

func main() {

	flag.Parse()

	filename := *file
	quizTime := *limit
	shuf := *shuffle

	quiz(filename, quizTime, shuf)
}
