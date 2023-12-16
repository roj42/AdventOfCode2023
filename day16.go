package main

import (
	"bufio"
	"fmt"
)

type symbolGrid [][]symbolLocation

type symbolLocation struct {
	symbol byte
	loc    coord
	visits int
}

func (sg symbolGrid) countVisits() int {
	total := 0
	for y := range sg {
		for x := range sg[y] {
			total += sg[y][x].visits
		}
	}
	return total
}

func (s symbolGrid) at(at coord) symbolLocation {
	if at.x < 0 || at.y < 0 || at.y+1 > len(s) || at.x+1 > len(s[at.y]) {
		return symbolLocation{}
	}
	return s[at.y][at.x]
}

func (s *symbolGrid) zap(inBeam beam) (newBeams []beam) {
	//update visited
	(*s)[inBeam.at.y][inBeam.at.x].visits++
	//swtich on symbol
	curSymbol := s.at(inBeam.at)
	switch curSymbol.symbol {
	case '.': // beam continues on unabated
		newCoord := navFrom(inBeam.at, inBeam.from)
		newBeams = append(newBeams, beam{at: newCoord, from: inBeam.from})
	case '|': //add two beams when ya should
	case '-': //add two beams when ya should

	}

	return newBeams
}

// a beam is the location where we are, and wherefrom it arrived.
type beam struct {
	at   coord
	from dir
}

// visualize a grid. Takes in an int, if that is a row number, it will visuzlize just that row
func (sg symbolGrid) visualize(row int) {
	startRow := 0
	endRow := len(sg)

	if row >= 0 && row < len(sg) {
		startRow = row
		endRow = row + 1
	}

	for i := startRow; i < endRow; i++ {
		visRow := []byte{}
		for _, sl := range sg[i] {
			visRow = append(visRow, sl.symbol)
		}
		fmt.Println(string(visRow))
	}
}

func day16(scanner *bufio.Scanner, isPart2 bool) string {

	grandTotal := 0

	//day 14 is one grid
	contraption := symbolGrid{}
	countY := 0
	for scanner.Scan() {
		//scan until blank line
		line := []symbolLocation{}
		for x, sym := range scanner.Text() {
			loc := symbolLocation{symbol: byte(sym), loc: coord{y: countY, x: x}}
			line = append(line, loc)
		}
		contraption = append(contraption, line)
		countY++
	}

	//let's go iterative, and set up some laser paths.

	beamsToTry := []beam{{at: coord{0, 0}, from: LEFT}}
	for i := 0; i < len(beamsToTry); i++ {
		newBeams := contraption.zap(beamsToTry[i])

		if len(newBeams) > 0 {
			beamsToTry = append(beamsToTry, newBeams...)
		}

		//what's this guy look like, and do we need to add more beams?
	}

	//did we error in there somewhere?
	if err := scanner.Err(); err != nil {
		check(err)
	}

	contraption.visualize(-1)
	log("visits", contraption.countVisits())

	if isPart2 {
		log("it sure is part 2")
	}

	return fmt.Sprint(grandTotal)
}
