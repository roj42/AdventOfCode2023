package main

import (
	"bufio"
	"fmt"
)

type symbolGrid [][]symbolLocation

type symbolLocation struct {
	symbol         byte
	loc            coord
	visits         int
	visitFromUP    bool
	visitFromDOWN  bool
	visitFromLEFT  bool
	visitFromRIGHT bool
}

func (s *symbolLocation) visit(d dir) bool {
	//mark that we were visited, and what direction.
	//if we already were, return true
	(*s).visits++
	switch d {
	case UP:
		cur := s.visitFromUP
		(*s).visitFromUP = true
		return cur
	case DOWN:
		cur := s.visitFromDOWN
		(*s).visitFromDOWN = true
		return cur
	case LEFT:
		cur := s.visitFromLEFT
		(*s).visitFromLEFT = true
		return cur
	case RIGHT:
		cur := s.visitFromRIGHT
		(*s).visitFromRIGHT = true
		return cur
	default:
		log("location visitited from an unknown direction")
		return false
	}
}

func (sg symbolGrid) countVisits() (int, int) {
	total := 0
	visited := 0
	for y := range sg {
		for x := range sg[y] {
			total += sg[y][x].visits
			if sg[y][x].visits > 0 {
				visited++
			}
		}
	}
	return total, visited
}

func (s symbolGrid) at(at coord) symbolLocation {
	if at.x < 0 || at.y < 0 || at.y+1 > len(s) || at.x+1 > len(s[at.y]) {
		return symbolLocation{symbol: byte(NOPE)}
	}
	return s[at.y][at.x]
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

	}

	//did we error in there somewhere?
	if err := scanner.Err(); err != nil {
		check(err)
	}

	// contraption.visualize(-1)
	total, grandTotal := contraption.countVisits()
	log("total visits", total)

	if isPart2 {
		log("it sure is part 2")
	}

	return fmt.Sprint(grandTotal)
}

func (s *symbolGrid) zap(inBeam beam) (newBeams []beam) {
	//swtich on symbol
	curSpot := s.at(inBeam.at)
	if curSpot.symbol == byte(NOPE) { //off the wall
		return
	}
	//update visited
	if (*s)[inBeam.at.y][inBeam.at.x].visit(inBeam.from) {
		log("loop at", inBeam.at.y, inBeam.at.x)
		return nil
	}

	prismDir := prismBounce(curSpot.symbol, inBeam.from)
	for _, pd := range prismDir {
		newCoord := navTo(inBeam.at, pd)
		newBeams = append(newBeams, beam{at: newCoord, from: op(pd)})
	}

	return newBeams
}

func prismBounce(prism byte, from dir) []dir {
	switch prism {
	case '.': //a dot will just keep going, exiting the opposite of from
		return []dir{op(from)}
	case '/': //mirror! up<-> left, down<->right
		switch from {
		case UP:
			return []dir{LEFT}
		case LEFT:
			return []dir{UP}
		case DOWN:
			return []dir{RIGHT}
		case RIGHT:
			return []dir{DOWN}
		}
	case '\\': //mirror! up<-> rught, down<->left
		switch from {
		case UP:
			return []dir{RIGHT}
		case RIGHT:
			return []dir{UP}
		case DOWN:
			return []dir{LEFT}
		case LEFT:
			return []dir{DOWN}
		}
	case '-': //splitter, double up/down, pass through left/right
		switch from {
		case UP:
			fallthrough
		case DOWN:
			return []dir{RIGHT, LEFT}
		case LEFT:
			fallthrough
		case RIGHT:
			return []dir{op(from)}
		}
	case '|': //splitter, double right/left, pass through up/down
		switch from {
		case UP:
			fallthrough
		case DOWN:
			return []dir{op(from)}
		case LEFT:
			fallthrough
		case RIGHT:
			return []dir{UP, DOWN}
		}
	}
	log("prismbounce did not recognize symbol", prism)
	return nil //notrecognized, no motion
}
