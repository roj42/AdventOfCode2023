package main

import (
	"bufio"
	"fmt"
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
		for combination := range GenerateCombinations(".#", qMarks) {
			if checkCombination(condition, combination, checks) {
				good += 1
			}
		}
		log(condition, checks, "\t\tgood:", good)
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
	// fmt.Print(".")
	//is it answer? let's make a checks for this combo:
	condchecks := []int{}
	countCurDam := 0
	for _, c := range condition {
		if c == '#' {
			countCurDam++
		} else { // dangerously assume a dot
			if countCurDam > 0 {
				condchecks = append(condchecks, countCurDam)
				countCurDam = 0
			}
		}
	}
	//maybe we hit the end.
	if countCurDam > 0 {
		condchecks = append(condchecks, countCurDam)
	}

	if binMask(condchecks) == checks {
		return true
	}

	return false
}

// a bit of borrowed code:
func GenerateCombinations(alphabet string, length int) <-chan string {
	c := make(chan string, 2)

	// Starting a separate goroutine that will create all the combinations,
	// feeding them to the channel c
	go func(c chan string) {
		defer close(c) // Once the iteration function is finished, we close the channel

		AddLetter(c, "", alphabet, length) // We start by feeding it an empty string
	}(c)

	return c // Return the channel to the calling function
}

// AddLetter adds a letter to the combination to create a new combination.
// This new combination is passed on to the channel before we call AddLetter once again
// to add yet another letter to the new combination in case length allows it
func AddLetter(c chan string, combo string, alphabet string, length int) {
	// Check if we reached the length limit
	// If so, we just return without adding anything
	if length <= 0 {
		c <- combo
		return
	}

	var newCombo string
	for _, ch := range alphabet {
		newCombo = combo + string(ch)
		AddLetter(c, newCombo, alphabet, length-1)
	}
}
