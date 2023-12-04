package main

import (
	"bufio"
	"fmt"
	"strconv"
)

type engineSchematic [][]byte

var schematic = engineSchematic{}

func day3(scanner *bufio.Scanner) string {
	//pre declare output at 0
	grandTotal := 0
	//Scanlines has some trouble with the default buffer.
	const maxCapacity int = 19750 // your required line length
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	//Let's load up our schematic as a matrix-o'-bytes
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		schematic = append(schematic, scanner.Bytes())
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

// parse
func processLineDay3(row int, line []byte) (rowTotal int) {
	log(string(line))
	advance := 0 //use this to skip forward.
	for col, char := range line {
		if advance > 0 {
			advance--
			continue
		}
		if isDigit(char) {
			pn := getPartNumberAt(row, col)
			advance = len(pn) - 1

			if isConnected(row, col, len(pn)) {
				log(pn, "connected")
				partValue, err := strconv.Atoi(pn)
				check(err)
				rowTotal += partValue
			} else {
				log(pn, " not connected")
			}
		}
	}
	return
}

func isDigit(b byte) bool {
	_, err := strconv.Atoi(string(b))
	return err == nil
}

func getPartNumberAt(row, col int) string {
	trimRow := schematic[row][col:]
	pn := ""
	for _, c := range trimRow {
		if isDigit(c) {
			pn += string(c)
		} else {
			break
		}
	}

	return pn
}

func isConnected(row, col, width int) bool {
	// check a rectangle one wider than our part number
	//but not, like, off the edges
	left := max(col-1, 0)
	right := min(col+width+1, len(schematic[0]))

	rowAbove := false
	if row-1 >= 0 {
		rowAbove = containsSymbol(schematic[row-1][left:right])
	}
	rowOn := false
	if !rowAbove {
		rowOn = containsSymbol(schematic[row][left:right])
	}
	rowBelow := false

	if row+1 != len(schematic) && !rowAbove {
		rowBelow = containsSymbol(schematic[row+1][left:right])

	}

	return rowAbove || rowOn || rowBelow
}

func containsSymbol(bytes []byte) bool {
	// log("checking :", string(bytes), " for symbols")
	check := false
	for _, c := range bytes {
		if !isDigit(c) && c != '.' {
			check = true
		}
	}
	// log("so... ", check)
	return check
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
