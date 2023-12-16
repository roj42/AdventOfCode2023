package main

import (
	"bufio"
	"fmt"
)

func day16(scanner *bufio.Scanner, isPart2 bool) string {

	grandTotal := 0
	//day 14 is one grid
	platform := platform{}
	for scanner.Scan() {
		//scan until blank line
		platform = append(platform, scanner.Text())
	}
	//did we error in there somewhere?
	if err := scanner.Err(); err != nil {
		check(err)
	}

	if isPart2 {
		log("it sure is part 2")
	}

	platform.tilt(UP)

	return fmt.Sprint(grandTotal)
}
