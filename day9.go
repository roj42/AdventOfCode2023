package main

import (
	"bufio"
	"fmt"
)

func day9(scanner *bufio.Scanner, isPart2 bool) string {

	//we'll get a grand total each row
	grandTotal := 0

	for scanner.Scan() {

		//read in a line
		//let's get arbitrarily large, baby
		workSpace := [][]int{}

		//We got ints
		workSpace = append(workSpace, toInts(scanner.Text()))
		//now calculate differences and append, loop until we get all zeroes
		for i := 0; ; i++ {
			newRow := []int{}
			for j := range workSpace[i] {
				if j == len(workSpace[i])-1 {
					break
				}
				newRow = append(newRow, (workSpace[i][j+1] - workSpace[i][j]))
			}
			workSpace = append(workSpace, newRow)
			countZeroes := countMatches(newRow, []int{0}) //reused code yay. A little inefficient
			if countZeroes == len(newRow) {
				break
			}
		}

		if isPart2 {
			grandTotal += analyzeWorkSpacePart2(workSpace)
		} else {
			grandTotal += analyzeWorkSpacePart1(workSpace)
		}

	}
	//did we error in there somewhere?
	if err := scanner.Err(); err != nil {
		check(err)
	}
	for row, line := range schematic {
		grandTotal += processLineDay3(row, line)
	}

	return fmt.Sprint(grandTotal)
}

func analyzeWorkSpacePart1(workSpace [][]int) (answer int) {
	//let's take that workspace and work backwards to apply the algorithm to the ends
	for i := len(workSpace) - 1; i > 0; i-- {
		//append on to the next row up: the last value of this row plus the last value of that row
		lastRow := i - 1
		lastEleLastRow := workSpace[lastRow][len(workSpace[lastRow])-1]
		lastEleThisRow := workSpace[i][len(workSpace[i])-1]
		workSpace[lastRow] = append(workSpace[lastRow], lastEleLastRow+lastEleThisRow)
	}
	answer = workSpace[0][len(workSpace[0])-1]
	log("line done, adding ", answer)
	return answer

}

func analyzeWorkSpacePart2(workSpace [][]int) (answer int) {
	//let's take that workspace and work backwards to apply the algorithm
	for i := len(workSpace) - 1; i > 0; i-- {
		//append on to the front of the next row up/prev row:
		// the first value of the prev row minus the first value of this one
		lastRow := i - 1
		firstElePrevRow := workSpace[lastRow][0]
		firstEleThisRow := workSpace[i][0]
		workSpace[lastRow] = append([]int{firstElePrevRow - firstEleThisRow}, workSpace[lastRow]...)
	}
	answer = workSpace[0][0]
	log("line done, adding ", answer)
	return answer

}
