package main

import (
	"bufio"
	"errors"
	"fmt"
)

type coord struct {
	y int
	x int
}

type diagram [][]byte

//this is fuckin' confusing. Let's make... translators, I guess.

type dir byte

func (d dir) getCoord() coord {
	switch d {
	case UP:
		return coord{y: -1}
	case DOWN:
		return coord{y: 1}
	case LEFT:
		return coord{x: -1}
	case RIGHT:
		return coord{x: 1}
	}
	return coord{}
}

const DOWN dir = 'd'
const UP dir = 'u'
const LEFT dir = 'l'
const RIGHT dir = 'r'
const NOPE dir = '!'

// op swapposite
func op(dr dir) dir {
	if dr == DOWN {
		return UP
	}
	if dr == UP {
		return DOWN
	}
	if dr == LEFT {
		return RIGHT
	}
	if dr == RIGHT {
		return LEFT
	}
	return NOPE
}

func (d diagram) connects(from dir, at coord) dir {
	junct := d.at(at) //The larger 4-legged one that assaulted Hoth.
	//at also kindly checks for length

	if junct != '!' {

		//what's smart? Nested switches, that's what
		switch from {
		case LEFT: //from the left
			switch junct {
			case '-':
				return RIGHT
			case 'J':
				return UP
			case '7':
				return DOWN
			}
		case RIGHT: //from the right
			switch junct {
			case '-':
				return LEFT
			case 'L':
				return UP
			case 'F':
				return DOWN

			}
		case DOWN: //from below
			switch junct {
			case '|':
				return UP
			case '7':
				return LEFT
			case 'F':
				return RIGHT
			}
		case UP: //from above
			switch junct {
			case '|':
				return DOWN
			case 'L':
				return RIGHT
			case 'J':
				return LEFT
			}
		}
	}

	log("bad input from", string(from), "to", at)
	return NOPE
}

// just dumb math for you
func navTo(at coord, to dir) coord {
	switch to {
	case UP:
		return coord{at.y - 1, at.x}
	case DOWN:
		return coord{at.y + 1, at.x}
	case LEFT:
		return coord{at.y, at.x - 1}
	case RIGHT:
		return coord{at.y, at.x + 1}
	}

	//return no movement if we couldn't move.
	log("bad move to", string(to), "at", at)
	return at
}

func (d diagram) at(at coord) byte {
	if at.x < 0 || at.y < 0 || at.y+1 > len(d) || at.x+1 > len(d[at.y]) {
		return '!'
	}
	return d[at.y][at.x]
}

func day10(scanner *bufio.Scanner, isPart2 bool) string {

	//scan in the entire map. we know it's 140x140, so that's small enought to just use a growing slice
	diag := diagram{}
	startPoint := coord{}
	//Scanlines has some trouble with the default buffer.
	const maxCapacity int = 19750 // your required line length
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	//Let's load up our schematic as a matrix-o'-bytes
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		line := scanner.Text()
		for i := range line {
			if line[i] == 'S' {
				startPoint.x = i
				startPoint.y = len(diag)
				break
			}
		}
		diag = append(diag, []byte(line))
	}
	//did we error in there somewhere?
	if err := scanner.Err(); err != nil {
		check(err)
	}

	//starting at startPoint

	//set next as 0,0, it happens to be safe for input

	//find a next from S
	next := coord{0, 0}
	log("starting with default next", next)
	var workingDir dir = NOPE
	for { //sloppy if else saver
		righty := navTo(startPoint, RIGHT)
		if c := diag.connects(LEFT, righty); c != NOPE {
			workingDir = LEFT
			next = righty
			break
		}
		lefty := navTo(startPoint, LEFT)
		if c := diag.connects(RIGHT, lefty); c != NOPE {
			next = lefty
			break
		}
		upso := navTo(startPoint, UP)
		if c := diag.connects(DOWN, upso); c != NOPE {
			next = upso
			break
		}
		downBaby := navTo(startPoint, DOWN)
		if c := diag.connects(UP, downBaby); c != NOPE {
			next = downBaby
			break
		}
		check(errors.New("no path from start"))
	}

	count := 1
	for ; diag.at(next) != 'S' && workingDir != NOPE; count++ {
		workingDir = diag.connects(workingDir, next)
		next = navTo(next, workingDir)
		workingDir = op(workingDir)
	}

	if !isPart2 {
		//all we need is the size of the route to calculate part 1
		log("we went", fmt.Sprint(count), "so half that")
		return fmt.Sprint(count / 2)
	}
	log("A path ", fmt.Sprint(count), "long")
	return "2"
}
