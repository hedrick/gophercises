package main

import (
	"encoding/csv"
	"io"
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

type question struct {
	Problem string
	Answer  string
}

func main() {
	readCSV("problems.csv")
}

// ReadCSV reads in a csv
// Returns a []Question and error if applicable
func readCSV(filename string) ([]question, error) {
	// open the file
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	questions := []question{}

	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			// update to sane error
			panic(err)
		}
		if len(line) != 2 {
			panic(err)
		}

		questions = append(questions, question{line[0], line[1]})
	}
	return questions, nil
}
