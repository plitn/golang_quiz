package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type problem struct {
	question string
	answer string
}

// lines = [[1+1 2] [2+2 4] [2+3 5]]
func linesToProoblems(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for idx, line := range lines{
		problems[idx] = problem {
			question: line[0],
			answer: line[1],
		}
	}
	return problems
}



func main() {
	fileName := flag.String("csv", "quiz.csv", "format: 'q,a'")
	tl := flag.Int("limit", 30, "quiz timelimit in seconds")
	flag.Parse()

	// делаем * потому что fileName Это указатель
	file, err := os.Open(*fileName)
	if err != nil {
		exit("failed to open file")
	}
	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()
	if err != nil {
		exit("some csv parse error")
	}

	timer := time.NewTimer(time.Duration(*tl) * time.Second)
	probs := linesToProoblems(lines)
	score := 0
	for idx, problem := range probs {
		fmt.Printf("Quiz question #%d: %s = ", idx + 1, problem.question)
		ansChan := make(chan string)
		go func() {
			var userAnswer string
			fmt.Scanln(&userAnswer)
			ansChan <- userAnswer
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nyou ran out ouf time and scored %d", score)
			return
		case userAnswer := <-ansChan:
			if userAnswer == problem.answer {
				score++
			}
		}
	}
	fmt.Printf("you scored %d out of %d", score, len(probs))
}



func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}