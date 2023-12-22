package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type problem struct{
	qustion string
	answer string
}

func main() {
	csvfilename := flag.String("csv", "problems.csv", "path to csv file")
	timeLimit := flag.Int("duration", 30, "number of seconds to spend during the quiz")
	flag.Parse()

	file, err := os.Open(*csvfilename)
	if err != nil{
		exit(err)
	}
	
	problems := readFile(file)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correntAnswer := 0
	loop:
	for i, p := range problems{
		fmt.Printf("Problem %d: %s = ", i+1, p.qustion)
		ansCh := make(chan string)
		go getAnswer(ansCh)

		select {
		case <-timer.C:
			fmt.Println()
			break loop
		case ans := <-ansCh:
			if p.answer == ans{
				correntAnswer++
			}
		}
	}

	fmt.Printf("Result %d out of %d", correntAnswer, len(problems))
}

func getAnswer(ansCh chan string){
	var answer string
	fmt.Scanf("%s\n", &answer)
	ansCh <- answer
}

func readFile(rd io.Reader)[]problem{
	r := csv.NewReader(rd)
	lines, err := r.ReadAll()
	if err != nil{
		exit(fmt.Errorf("failed to read cvs contents: %v", err))
	}
	return parselines(lines)
}

func parselines(lines [][]string) []problem{
	arr := make([]problem, len(lines))
	for i, line := range lines {
		arr[i] = problem{
			qustion: line[0],
			answer: strings.TrimSpace(line[1]),
		}
	}
	return arr
}

func exit(err error){
	fmt.Printf("some error occurred: %v\n", err)
	os.Exit(1)
}