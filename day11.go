package main

import (
	"bufio"
	"fmt"
	"math"
	"slices"
)

// we can reuse universeram, but then we'd never have a STARMAP
type starmap [][]byte

var EXPAND_O_FACTOR = 1

func day11(scanner *bufio.Scanner, isPart2 bool) string {
	if isPart2 {
		EXPAND_O_FACTOR = 999999
	}

	//scan in the entire star map. we know it's 140x140, so that's small enough to just use a growing slice
	universe := starmap{}
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		line := scanner.Text()
		universe = append(universe, []byte(line))
	}
	//did we error in there somewhere?
	if err := scanner.Err(); err != nil {
		check(err)
	}
	log("starting universe is ", len(universe), "by", len(universe[0]))
	//GALACTIC EXPLORATION
	clearRows := []int{}
	for i := range universe {
		if universe.rowClearOfGalaxies(i) {
			clearRows = append(clearRows, i)
		}
	}
	log("found clear rows:", len(clearRows))
	// universe.addRowsAfter(clearRows)

	clearCols := []int{}
	for i := range universe[0] {
		if universe.colClearofGalaxies(i) {
			clearCols = append(clearCols, i)
		}
	}
	log("found clear columns:", len(clearCols))
	// universe.addColsAfter(clearCols)

	log("expanded universe is ", len(universe)+len(clearRows), "by", len(universe[0])+len(clearCols))

	//let's find galaxies
	galCoords := findSymbols(universe, '#')

	grandTotal := 0
	//min dist in steps is just a diff of x and y. Double loop, skipping a dif with yourself
	neighCoords := append([]coord{}, galCoords...)
	for _, curCoord := range galCoords {
		neighCoords = neighCoords[1:]
		for _, neigCoord := range neighCoords {
			grandTotal += int(math.Abs(float64(curCoord.x)-float64(neigCoord.x)) + math.Abs(float64(curCoord.y)-float64(neigCoord.y)))
			//let's see if there are any rows or columns in between. If so, add the EXPAND_O_FACTOR
			for _, val := range clearRows {
				from := min(curCoord.y, neigCoord.y)
				to := max(curCoord.y, neigCoord.y)
				if from <= val && val <= to {
					grandTotal += EXPAND_O_FACTOR
				}
			}
			for _, val := range clearCols {
				from := min(curCoord.x, neigCoord.x)
				to := max(curCoord.x, neigCoord.x)
				if from <= val && val <= to {
					grandTotal += EXPAND_O_FACTOR
				}
			}
		}
	}

	return fmt.Sprint(grandTotal)
}

func findSymbols(grid [][]byte, sym byte) (coordList []coord) {
	for y, row := range grid {
		for x, char := range row {
			if char == sym {
				coordList = append(coordList, coord{y: y, x: x})
			}
		}
	}
	return
}

func (s starmap) vis() {
	for _, r := range s {
		log(string(r))
	}
}

func (s starmap) colClearofGalaxies(colX int) bool {
	for rowY := range s {
		if s.at(coord{y: rowY, x: colX}) != '.' {
			return false
		}
	}
	return true
}

// expands a startmap vertically by doubling rows given
func (s *starmap) addColsAfter(cols []int) {
	//let's make sure we're in reverse order here:
	slices.Sort(cols)
	slices.Reverse(cols)
	for row := range *s {
		for _, c := range cols {
			//double the row in one append, by appending 0toX with XtoEnd, two X-es
			new := append((*s)[row][:c+1], (*s)[row][c:]...)
			(*s)[row] = new
		}
	}

}

func (s starmap) rowClearOfGalaxies(rowY int) bool {
	for colX := range s[rowY] {
		if s.at(coord{y: rowY, x: colX}) != '.' {
			return false
		}
	}
	return true
}

// expands a startmap vertically by doubling rows given
func (s *starmap) addRowsAfter(rows []int) {
	//let's make sure we're in reverse order here:
	slices.Sort(rows)
	slices.Reverse(rows)
	for _, r := range rows {
		new := append((*s)[:r+1], (*s)[r:]...)
		*s = new
	}
}

// at gives you what's there, unless it's off the map, then it saves you.
func (s starmap) at(at coord) byte {
	if at.x < 0 || at.y < 0 || at.y+1 > len(s) || at.x+1 > len(s[at.y]) {
		return '!'
	}
	return s[at.y][at.x]
}
