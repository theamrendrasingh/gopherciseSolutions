package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func quiz(filename string, quizTime int, shuf bool) {

	quizRecord := csvReader(filename)

	if shuf == true {
		quizRecord = shuffleQuestions(quizRecord)
	}

	correctCount := 0

	ansChan := make(chan string)

	ticker := time.NewTicker(time.Duration(quizTime) * time.Second)

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
			if cleanString(input) == cleanString(ans) {
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

func cleanString(s string) string {
	s = strings.TrimRight(s, "\n")
	s = strings.ToLower(s)
	return s
}

func shuffleQuestions(slice [][]string) [][]string {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	for n := len(slice); n > 0; n-- {
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
	}
	return slice
}
