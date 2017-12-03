package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/rbickham11/adventofcode/days"
)

var problemFunctions = map[int]func() (string, string){
	1: days.One,
	2: days.Two,
}

func main() {
	flag.Parse()

	problemNumberArg := flag.Arg(0)

	if len(problemNumberArg) == 0 {
		fmt.Println("Usage: \"go run main {day_number}\"")
		os.Exit(1)
	}

	problemNumber, err := strconv.Atoi(flag.Arg(0))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if problemFunction, ok := problemFunctions[problemNumber]; ok {
		answerA, answerB := problemFunction()
		fmt.Printf("Day %d: a: %s b: %s\n", problemNumber, answerA, answerB)
		os.Exit(0)
	}

	if problemNumber <= 25 {
		fmt.Printf("Puzzle for day %d has not been solved yet!\n", problemNumber)
		os.Exit(1)
	}

	fmt.Println("Please enter a valid day!")
}
