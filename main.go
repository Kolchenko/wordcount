package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	input := readInput()
	fmt.Println(count(input))
}

func readInput() (input string) {
	flag.Parse()
	input = strings.Join(flag.Args(), "")
	return input
}

func count(str string) int {
	return len(strings.Fields(str))
}
