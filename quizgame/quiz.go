package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type Problem struct {
	question      string
	answerOptions []string
}

func exitOnError(errorMsg string, errCode int) {
	fmt.Println(errorMsg)
	os.Exit(errCode)
}

func linesToProblemSlice(lines [][]string) []Problem {
	result := make([]Problem, len(lines))
	for i, line := range lines {
		result[i] = Problem{question: line[0], answerOptions: line[1:]}
	}
	return result
}

func isAnswerCorrect(answer string, pr Problem) bool {
	for _, ansOpt := range pr.answerOptions {
		if ansOpt == answer {
			return true
		}
	}
	return false
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "Path to csv file containing problems")
	timeLimit := flag.Int("timelimit", 10, "Time limit to answer the question. If the timer expires, you are going to answer next qustion")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exitOnError("Failed to open file "+*csvFilename, 1)
	}
	csvReader := csv.NewReader(file)
	lines, err := csvReader.ReadAll()
	if err != nil {
		exitOnError("Can't parse csv file", 1)
	}
	problems := linesToProblemSlice(lines)
	var correctAnswersCounter int
	for i, pr := range problems {
		fmt.Printf("Problem №%d: %s\nYour answer: ", i+1, pr.question)
		timer := time.NewTimer(time.Duration(*timeLimit * int(time.Second)))
		timerFired := false
		go func() {
			<-timer.C
			fmt.Println("Your time is up. Going to the next question. Press enter to show the next question")
			timerFired = true
		}()
		var answer string
		fmt.Scanf("%s\n", &answer)
		if !timerFired {
			timer.Stop()
			if isAnswerCorrect(answer, pr) {
				fmt.Println("You are correct!")
				correctAnswersCounter++
			} else {
				fmt.Println("Unfortunately, you are not correct")
			}
		}
	}
	fmt.Printf("You answered %d correct questions out of %d. It's %f percent", correctAnswersCounter, len(problems), float32(correctAnswersCounter)/float32(len(problems))*100)
}
