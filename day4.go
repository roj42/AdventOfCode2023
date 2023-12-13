package main

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func day4(scanner *bufio.Scanner) string {

	//pre declare output at 0
	grandTotal := 0

	for scanner.Scan() {
		log("line:" + scanner.Text())

		//Line step: parse each line for sweet sweet details
		numWinners := processLineDay4(scanner.Text())
		log(fmt.Sprint("parsed Line: ", numWinners))

		//Aggregation step: do we add? what are we adding?
		//the value of the line doubles as numWinners goes up
		worth := int(math.Pow(float64(2), float64(numWinners-1)))
		log("Line worth: ", worth)
		grandTotal += worth

	}

	//scanner is weird about errors. It will kick us out of the loop that .Scan() produces when there is one, so do this odd check.
	if err := scanner.Err(); err != nil {
		check(err)
	}

	return fmt.Sprint(grandTotal)
}

func processLineDay4(line string) (numWinners int) {
	//lets parse it out into two lists
	winners, available := splitListsDay4(line)

	//how many match
	numWinners = countMatches(winners, available)
	return
}

func splitListsDay4(line string) (winners, available []int) {
	//let's get left and right of the bar
	halves := strings.Split(line, "|")
	//let's get the gamestring out of the left side
	gameAndLeft := strings.Split(halves[0], ":")
	halves[0] = gameAndLeft[1]
	//now let's turn these into lists of numbers
	winners = toInts(strings.Trim(halves[0], " "), " ")
	available = toInts(strings.Trim(halves[1], " "), " ")
	return
}

func toInts(stringOfInts, splitOn string) []int {
	listOfInts := []int{}
	for _, str := range strings.Split(stringOfInts, splitOn) {
		//note that we might have bonus spaces for single digit numbers, so skip those
		if str == "" {
			continue
		}
		parsedInt, e := strconv.Atoi(str)
		check(e)
		listOfInts = append(listOfInts, parsedInt)
	}
	return listOfInts
}

// return the number times attempts has a number that matches a number in hits
func countMatches(hits, attempts []int) int {
	matches := 0
	for _, hit := range hits {
		for _, attempt := range attempts {
			if attempt == hit {
				matches += 1
				break
			}
		}
	}
	return matches
}

func day4_2(scanner *bufio.Scanner) string {

	//pre declare output at 0
	grandTotal := 0

	//make a map to count up cards
	cardsAndCounts := make(map[int]int)

	cardID := 1
	for scanner.Scan() {
		log("line:" + scanner.Text())
		//add 1 to the count of this map id
		curCount := cardsAndCounts[cardID]
		cardsAndCounts[cardID] = curCount + 1

		//Line step: parse each line for sweet sweet details
		numWinners := processLineDay4(scanner.Text())
		log(fmt.Sprint("parsed Line value: ", numWinners))
		//Add to counts of the next numWinners cards. You should add the number of cards in this group, since each copy will add 1
		for numWinners > 0 {
			curCount := cardsAndCounts[cardID+numWinners]
			cardsAndCounts[cardID+numWinners] = curCount + cardsAndCounts[cardID]
			numWinners--
		}
		cardID++

	}
	//scanner is weird about errors. It will kick us out of the loop that .Scan() produces when there is one, so do this odd check.
	if err := scanner.Err(); err != nil {
		check(err)
	}

	//Aggregation step: do we add? what are we adding?
	//just add up the counts
	for _, k := range cardsAndCounts {
		grandTotal += k
	}

	return fmt.Sprint(grandTotal)
}
