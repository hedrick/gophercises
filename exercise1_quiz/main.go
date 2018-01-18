package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

// Read in a file (csv) ✅
// Present the quiz to the user
// Keep track of how many questions they get right and how many incorrect
// Immediately ask the next question whether the previous is wrong/right
// Default to problems.csv but user should be able to customize the file via
// a flag

// CSV format:
// 5+5, 10
// 7+3, 10
// 1+1, 2
// etc

// Assume quizzes are < 100 questions and have single word/number answers

// At the end, should output total number of questions correct
// and how many questions there were in total (x/y correct)
// Invalid answers are considered incorrect.

func main() {

	csvFileName := flag.String("csv", "problems.csv",
		"a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV file: %s", *csvFileName))
		os.Exit(1)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Couldn't parse the CSV file."))
	}
	problems := parseLines(lines)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
