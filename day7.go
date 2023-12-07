package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"slices"
	"sort"
	"strings"
)

type handAndBid struct {
	Hand  string
	Bid   int
	Score int64
}

func day7(scanner *bufio.Scanner) string {

	handData := []handAndBid{}
	//read in them hands
	for scanner.Scan() {
		log("line:" + scanner.Text())
		input := strings.Split(scanner.Text(), " ")
		newHand := handAndBid{
			Hand: input[0],
			Bid:  toInts(input[1])[0],
		}
		//score it up. This is where savings will happen if you need 'em
		scoreHand(&newHand)
		handData = append(handData, newHand)

	}
	//scanner is weird about errors. It will kick us out of the loop that .Scan() produces when there is one, so do this odd check.
	if err := scanner.Err(); err != nil {
		check(err)
	}

	//sort hands by score
	sort.Slice(handData, func(i, j int) bool {
		return handData[i].Score < handData[j].Score
	})

	grandTotal := 0
	//now multiply bid by rank
	for rank, hand := range handData {
		grandTotal += (rank + 1) * hand.Bid
	}

	return fmt.Sprint(grandTotal)
}

const cardSymbolMask string = "23456789TJQKA"

// scorehand adds a score to a give hand in-place
func scoreHand(hnb *handAndBid) {
	//scoreing is hand value and then highest first card.

	//"base" score:
	// sum each 'face value' where 2 is 1 point, and ace is 13
	//times
	//10 to the power of double its position (reverse order) (1 is 100000000, 5 is 1)
	//If we do it this way, each two digits represents the face, and the multiplier ensures that
	//position is always worth more than value
	//ex 32222 = 200,000,000 + 1,000,000 + 10,000 + 100 + 1 = 201,010,101
	//ex "T55J5" = 900000000 + 4000000 + 40000 + 1000 + 4 = 904041004 or 09.04.04.10.04
	//It's math instead of symbols, so we can just do number sorts.

	handMap := map[rune]int{}
	for i, symbol := range hnb.Hand {
		//index 0 is worth 5, 4 is worth 1
		exp := float64(2 * (4 - i)) //we want each number to have two digits to occupy
		positionMult := int64(math.Pow(float64(10), exp))
		faceValue := strings.IndexRune(cardSymbolMask, symbol) + 1
		hnb.Score += int64(faceValue) * positionMult

		//For later, let's put the symbols in a map so we can calculate the hand
		c := handMap[symbol]
		handMap[symbol] = c + 1
	}

	//Score the actual hand.
	//the map will have the symbols and their counts. We don't care about the symbols, just the counts.
	//lets convert them to a code
	codeList := []string{}
	for _, count := range handMap {
		codeList = append(codeList, fmt.Sprint(count))
	}
	slices.Sort(codeList)
	slices.Reverse(codeList)
	code := string(strings.Join(codeList, ""))

	handScore := 0

	switch code {
	case "5": //5 of a kind
		handScore = 6
	case "41": //4 of a kind
		handScore = 5
	case "32": //everywhere you look
		handScore = 4
	case "311": //three of a kind
		handScore = 3
	case "221": //two pair
		handScore = 2
	case "2111": //one pair
		handScore = 1
	case "11111": //high card
		handScore = 0
	default:
		check(errors.New("hand code not recognized: " + code))
	}
	//There are 6 hands, and hand value is worth more than face/position, so multiply by 10^10
	hnb.Score = hnb.Score + (int64(handScore) * 10000000000)

}
