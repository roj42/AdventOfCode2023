package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"

	// "slices"
	"sort"
	"strings"
)

type handAndBid struct {
	Hand  string
	Bid   int
	Score int64
}

func day7(scanner *bufio.Scanner, isPart2 bool) string {
	if isPart2 {
		cardSymbolMask = jokersWildMask
	}

	handData := []handAndBid{}
	//read in them hands
	for scanner.Scan() {
		log("line:" + scanner.Text())
		input := strings.Split(scanner.Text(), " ")
		newHand := handAndBid{
			Hand: input[0],
			Bid:  toInts(input[1], " ")[0],
		}
		//score it up. This is where savings will happen if you need 'em
		scoreHand(&newHand, isPart2)
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

	log("Ad 'em up!")
	grandTotal := 0
	//now multiply bid by rank
	for rank, hand := range handData {
		log("rank:", rank, "hand: ", hand.Hand)
		grandTotal += (rank + 1) * hand.Bid
	}

	return fmt.Sprint(grandTotal)
}

var cardSymbolMask string = "23456789TJQKA"
var jokersWildMask string = "J23456789TQKA"

// scorehand adds a score to a give hand in-place
func scoreHand(hnb *handAndBid, jokersWild bool) {
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

	//facescore updates the face score of each and returns a handy map of their counts
	handMap, highCard := faceScore(hnb)

	if jokersWild && strings.Contains(hnb.Hand, "J") && len(handMap) > 1 { //well, fuck
		// move the joker count to the card that'd make the nicest hand.
		//add it to the highest count, unless equal, then add it to the highest face
		/*
				Highest card:
			5:	1111
			4:	111
			3:	22
			2:	1-
			2:	4-
				numerous card
			3:	31
			3:	21
			4:	211
		*/

		//remember and remove j
		jackCount := handMap['J']
		delete(handMap, 'J')
		//get the code of the rest and switch on it
		switch handMaptoCode(handMap) {
		case "1111":
			fallthrough
		case "111":
			fallthrough
		case "22":
			fallthrough
		case "1":
			fallthrough
		case "2":
			fallthrough
		case "3":
			fallthrough
		case "4":
			//these all should have the jack added to the most numerous card
			handMap[highCard] = handMap[highCard] + jackCount
		case "31":
			fallthrough
		case "21":
			fallthrough
		case "211":
			//these should go to the most numerous card
			count := 0
			numerousCard := '%'
			for k, val := range handMap {
				if val > count {
					numerousCard = k
					count = val
				}
			}
			handMap[numerousCard] = handMap[numerousCard] + jackCount
		default:
			check(errors.New("hand code (jack) not recognized: " + handMaptoCode(handMap) + ", hand: " + hnb.Hand))

		}

	}

	//Score the actual hand.
	handScore(hnb, handMap)

}

func faceScore(hnb *handAndBid) (map[rune]int, rune) {

	highCard := '%'
	highVal := -1
	handMap := map[rune]int{}

	for i, symbol := range hnb.Hand {
		//index 0 is worth 5, 4 is worth 1
		exp := float64(2 * (4 - i)) //we want each number to have two digits to occupy
		positionMult := int64(math.Pow(float64(10), exp))
		faceValue := strings.IndexRune(cardSymbolMask, symbol)
		//record highest card seen
		if faceValue > highVal {
			highCard = symbol
			highVal = faceValue
		}
		hnb.Score += (int64(faceValue) + 1) * positionMult
		//For later, let's put the symbols in a map so we can calculate the hand
		c := handMap[symbol]
		handMap[symbol] = c + 1
	}
	return handMap, highCard
}

func handScore(hnb *handAndBid, handMap map[rune]int) {
	//the map will have the symbols and their counts. We don't care about the symbols, just the counts.
	//lets convert them to a code
	code := handMaptoCode(handMap)

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
		check(errors.New("hand code not recognized: " + code + ", hand: " + hnb.Hand))
	}
	//There are 6 hands, and hand value is worth more than face/position, so multiply by 10^10
	hnb.Score = hnb.Score + (int64(handScore) * 10000000000)
}

func handMaptoCode(handMap map[rune]int) string {
	codeList := []string{}
	for _, count := range handMap {
		codeList = append(codeList, fmt.Sprint(count))
	}

	sort.Strings(codeList)
	last := len(codeList) - 1
	for i := 0; i < len(codeList)/2; i++ {
		codeList[i], codeList[last-i] = codeList[last-i], codeList[i]
	}
	return string(strings.Join(codeList, ""))
}
