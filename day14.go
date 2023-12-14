package main

import (
	"bufio"
	"fmt"
)

const ROCK = 'O'
const CUBE = '#'
const EMPT = '.'

type platform []string

func day14(scanner *bufio.Scanner, isPart2 bool) string {
	if isPart2 {
		log("it sure is part 2")

	}
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

	platform.tilt(UP)

	grandTotal += platform.weigh(UP)

	return fmt.Sprint(grandTotal)
}

func (p platform) weigh(direction dir) (totalWeight int) {
	//weight is platform dimension - index, so something right next to (0) the top, will weight <height> or 10
	for y, line := range p {
		for _, spot := range line {
			if spot == ROCK {
				switch direction {
				//if it's up, weight is platform height minus y dimension of the rock
				case UP:
					totalWeight += len(p) - y
				}
			}
		}
	}

	return totalWeight
}

func (p *platform) tilt(direction dir) {

	//Hmm, we need to range differently for other direcitons. here is up
	for rowI, line := range *p {
		for colI := range line {
			if line[colI] == ROCK {
				rockRow := rowI
				rockCol := colI
				for p.canMove(rockRow, rockCol, direction) {
					p.move(rockRow, rockCol, direction)
					coord := direction.getCoord()
					rockRow = rockRow + coord.y
					rockCol = rockCol + coord.x
				}
			}
		}
	}
}

func (p platform) canMove(rowI, colI int, UP dir) bool {
	coord := UP.getCoord()
	yMove := rowI + coord.y
	xMove := colI + coord.x
	//first, will we be w/in the platform still?
	if yMove >= 0 && yMove < len(p) && xMove >= 0 && xMove < len(p[rowI]) {
		return p[yMove][xMove] == EMPT
	}
	return false
}

func (p *platform) move(rowI, colI int, UP dir) {
	coord := UP.getCoord()
	yMove := rowI + coord.y
	xMove := colI + coord.x
	//update the two points; strings are immutable in go
	moveLine := []byte((*p)[yMove])
	moveLine[xMove] = (*p)[rowI][colI]
	(*p)[yMove] = string(moveLine)

	emptLine := []byte((*p)[rowI])
	emptLine[colI] = EMPT
	(*p)[rowI] = string(emptLine)
}
