package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func quiz(quizRecord [][]string) {

	correctCount := 0

	reader := bufio.NewReader(os.Stdin)

	for id, quesAns := range quizRecord {
		ques := quesAns[0]
		ans := quesAns[1]
		fmt.Println("Question ", id, " : ", ques)
		input, _ := reader.ReadString('\n')

		if strings.TrimRight(input, "\n") == ans {
			correctCount++
		}
	}

	fmt.Println("Questions Correctly Answered : ", correctCount)
}
