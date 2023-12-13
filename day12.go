package main

import (
	"bufio"
	"fmt"
	"strings"
)

// we can reuse universeram, but then we'd never have a STARMAP

func day12(scanner *bufio.Scanner, isPart2 bool) string {
	if isPart2 {
		log("it sure is part 2")
	}

	//day 12 is more a of a line by line style
	for scanner.Scan() {
		condition, checks := splitLine12(scanner.Text())
		log(condition, checks)
	}
	//did we error in there somewhere?
	if err := scanner.Err(); err != nil {
		check(err)
	}

	grandTotal := 0

	return fmt.Sprint(grandTotal)
}

func splitLine12(input string) (string, []int) {
	parts := strings.Split(input, " ")

	return parts[0], toInts(parts[1], ",")
}
