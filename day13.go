package main

import (
	"bufio"
	"fmt"
)

func day13(scanner *bufio.Scanner, isPart2 bool) string {
	if isPart2 {
		log("it sure is part 2")

	}
	grandTotal := 0
	//day 13 is chunks of data
	grids := [][]string{}
	curGrid := []string{}
	for scanner.Scan() {
		//scan until blank line
		line := scanner.Text()
		if len(line) == 0 {
			grids = append(grids, curGrid)
			curGrid = []string{}
			continue

		} else {
			curGrid = append(curGrid, line)
		}
	}
	//tack on the last one
	grids = append(grids, curGrid)

	for _, grid := range grids {
		rowsAbove := findHoriMirror(grid)
		if rowsAbove > 0 {
			grandTotal += (100 * rowsAbove)
		}
		colsLeft := findVertMirror(grid)
		if colsLeft > 0 {
			grandTotal += (colsLeft)
		}
	}

	//did we error in there somewhere?
	if err := scanner.Err(); err != nil {
		check(err)
	}

	return fmt.Sprint(grandTotal)
}

func findVertMirror(grid []string) int {
	//pivot
	return findHoriMirror(rotateGrid(grid))
}

// more mirroted then rotated, but, whatever.
func rotateGrid(grid []string) []string {
	tated := []string{}
	//assume a rectangular grid
	for i := range grid[0] {
		newString := ""
		for _, line := range grid {
			newString = newString + string(line[i])
		}
		tated = append(tated, newString)
	}
	return tated

}

func findHoriMirror(grid []string) int {
	//start at second row, and check above and below
	for curRow := range grid {
		if curRow == 0 {
			continue
		}
		//if the row and the previous row are equivalent, this is a possible mirror
		if grid[curRow] == grid[curRow-1] {
			mirrorRowsAbove := curRow
			upI := curRow - 2
			for downI := curRow + 1; ; downI++ {
				//end check, are we off the grid
				if upI < 0 || downI >= len(grid) {
					return mirrorRowsAbove
				}
				if grid[upI] != grid[downI] { //well that's not it
					break
				}
				upI--
			}
		}
	}
	return 0
}
