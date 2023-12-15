package main

import (
	"bufio"
	"fmt"
	"strings"
)

func day15(scanner *bufio.Scanner, isPart2 bool) string {

	//day 15 is one giant string
	steps := []string{}
	for scanner.Scan() {
		steps = strings.Split(scanner.Text(), ",")
	}

	//answer is the sum of hashes
	grandTotal := 0
	for _, step := range steps {
		grandTotal += xmasHash(step)
	}
	//did we error in there somewhere?
	if err := scanner.Err(); err != nil {
		check(err)
	}

	if isPart2 {
		log("it sure is part 2")
	}

	return fmt.Sprint(grandTotal)
}

func xmasHash(step string) int {
	hash := 0
	for _, char := range step {
		hash += int(char)
		hash = (hash * 17) % 256
	}
	log("hash of", step, "is", hash)
	return hash
}
