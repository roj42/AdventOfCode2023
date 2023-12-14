package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"strings"
)

func day12(scanner *bufio.Scanner, isPart2 bool) string {
	copyTimes := 0
	if isPart2 {
		log("it sure is part 2")
		copyTimes = 5
	}

	grandTotal := 0
	//day 12 is more of a line by line style
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		condition := parts[0]
		checkParts := toInts(parts[1], ",")

		//copy 5 times
		smCon := condition
		smChe := checkParts
		for i := 0; i < copyTimes; i++ {
			condition = condition + "?" + smCon
			checkParts = append(checkParts, smChe...)
		}

		checks := binMask(checkParts)
		// try every combo of working or not, then test it to see if it's okay

		//count up question marks
		qMarks := strings.Count(condition, "?")

		//generate combos
		good := 0
		for combination := range GenerateCombinations(qMarks) {
			if checkCombination(condition, combination, checks) {
				good += 1
			}
		}
		log(condition, checkParts, "good:", good)
		grandTotal += good
	}

	//did we error in there somewhere?
	if err := scanner.Err(); err != nil {
		check(err)
	}

	return fmt.Sprint(grandTotal)
}

func binMask(checks []int) (binChecks int) {
	for v, i := range checks {
		for c := 0; c < i; c++ {
			binChecks = binChecks << 1
			binChecks++
		}
		if v != len(checks)-1 {
			binChecks = binChecks << 1
		}
	}
	return
}

// create a nice new string, by swapping combination's charaters for conditions, and then compare to check.
func checkCombination(condition, combination string, checks int) bool {

	//build out the replacement, swapping characters in combination for ? in condition
	for _, c := range combination {
		//replace the first (remaining) question mark
		condition = strings.Replace(condition, "?", string(c), 1)
	}
	//is it answer? let's make a checks for this combo:
	condChecks := 0
	wasNotHash := false

	for v, c := range condition {
		if c == '#' {
			if wasNotHash {
				condChecks = condChecks << 1
			}
			condChecks++
			if v < len(condition)-1 {
				condChecks = condChecks << 1
			}
			wasNotHash = false
		} else {
			wasNotHash = true
			if v == len(condition)-1 { //one too many crank overs, and we're ending with a dot (or run of dots)
				condChecks = condChecks >> 1
			}
		}
	}

	return condChecks == checks
}

func GenerateCombinations(length int) <-chan string {
	c := make(chan string)
	// Starting a separate goroutine that will create all the combinations,
	// feeding them to the channel c
	go func(c chan string) {
		defer close(c) // Once the iteration function is finished, we close the channel
		for i := 1; bits.Len(uint(i)) <= length; i++ {
			prePad := reverseMask(i)
			postPad := PadLeft(prePad, ".", length)
			c <- postPad
		}
	}(c)
	return c
}

func reverseMask(i int) string {
	bytey := []byte{}
	for i > 0 {
		if i%2 == 0 { //is a zero, therefore .
			bytey = append([]byte{'.'}, bytey...)
		} else {
			bytey = append([]byte{'#'}, bytey...)
		}
		i = i >> 1
	}
	return string(bytey)
}

func PadLeft(str, pad string, lenght int) string {
	for {
		str = pad + str
		if len(str) > lenght {
			return str[1:]
		}
	}
}
