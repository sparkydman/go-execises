package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct{
	qustion string
	answer string
}

func main() {
	csvfilename := flag.String("csv", "problems.csv", "path to csv file")
	flag.Parse()

	file, err := os.Open(*csvfilename)
	if err != nil{
		exit(err)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil{
		exit(fmt.Errorf("failed to read cvs contents: %v", err))
	}
	problems := parselines(lines)
	
	correntAnswer := 0
	for i, p := range problems{
		fmt.Printf("Problem %d: %s =\n", i+1, p.qustion)
		var answer string
		fmt.Scanf("%s\n", &answer)

		if p.answer == answer{
			correntAnswer++
		}
	}

	fmt.Printf("Result %d out of %d", correntAnswer, len(problems))
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