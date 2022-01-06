package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := readInput()
	if err != nil {
		fail(err)
	}
	fmt.Println(count(input))
}

func readInput() (input string, err error) {
	flag.Parse()
	input = strings.Join(flag.Args(), "")
	if input == "" {
		return input, errors.New("missing string to count words")
	}
	return input, nil
}

func fail(err error) {
	fmt.Println("wordcount:", err)
	os.Exit(1)
}

func count(str string) int {
	return len(strings.Fields(str))
}
