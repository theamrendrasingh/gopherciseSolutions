package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func quiz(quizRecord [][]string, quizTime int) {

	correctCount := 0

	ansChan := make(chan string)

	fmt.Println("About to Start Timer")

	ticker := time.NewTicker(time.Duration(quizTime) * time.Second)

	fmt.Println("timer Started")

	for id, quesAns := range quizRecord {
		ques := quesAns[0]
		ans := quesAns[1]
		fmt.Println("Question ", id, " : ", ques)
		go getAnswer(ansChan)

		select {
		case <-ticker.C:
			fmt.Println("\nTime Over")
			fmt.Println("Questions Correctly Answered : ", correctCount)
			return
		case input := <-ansChan:
			if strings.TrimRight(input, "\n") == ans {
				correctCount++
			}
			break
		}
	}

	fmt.Println("Questions Correctly Answered : ", correctCount)
}

func getAnswer(ansChan chan string) {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	ansChan <- input
}
