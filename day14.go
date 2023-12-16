package main

import (
	"bufio"
	"fmt"
	"sort"
)

const ROCK = 'O'
const CUBE = '#'
const EMPT = '.'

type platform []string

func day14(scanner *bufio.Scanner, isPart2 bool) string {

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
		//start cyclin
		snapShots := []string{}
		snapShot := ""
		for _, line := range platform {
			snapShot = snapShot + line
		}
		snapShots = append(snapShots, snapShot)

		for i := 0; i < 1000; i++ {

			platform.tilt(UP)
			platform.tilt(LEFT)
			platform.tilt(DOWN)
			platform.tilt(RIGHT)
			snapShot := ""
			for _, line := range platform {
				snapShot = snapShot + line
			}
			if i%10000 == 0 {
				for j, ss := range snapShots {
					if snapShot == ss {
						log("cycle", i, "is the same as", j)
						// break
						return fmt.Sprint(platform.weigh(UP))
					}
				}
			}
			snapShots = append(snapShots, snapShot)
			if len(snapShots) > 1000 {
				snapShots = snapShots[:1000]
			}
		}
	} else {
		platform.tilt(UP)
	}

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

func (p platform) tilt(direction dir) {

	//make a list of indexes we can reverse.
	rowOrder := []int{}
	for index := range p {
		rowOrder = append(rowOrder, index)
	}
	if direction == DOWN {
		sort.Sort(sort.Reverse(sort.IntSlice(rowOrder)))
	}

	for _, rowI := range rowOrder {
		line := p[rowI]
		//make a list of indexes for line we can reverse
		lineOrder := []int{}
		for index := range line {
			lineOrder = append(lineOrder, index)
		}
		if direction == RIGHT {
			sort.Sort(sort.Reverse(sort.IntSlice(lineOrder)))
		}
		for _, colI := range lineOrder {
			//if we're going right, we need to reverse line

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
