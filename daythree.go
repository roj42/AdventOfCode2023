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
	for col := range line {
		if advance > 0 {
			advance--
			continue
		}

		pn := getPartNumberAt(row, col)
		if pn != "" {
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
	//there are no part numbers off the grid
	if row < 0 || row >= len(schematic) || col < 0 || col >= len(schematic[0]) {
		return ""
	}

	//if we don't have a digit, we don't have a PN, so piss off
	if !isDigit(schematic[row][col]) {
		return ""
	}

	//go backwards until we get a non-digit
	for i := col; i >= 0; i-- {
		//if it's not a digit, update col to indicate the start of the part number
		if !isDigit(schematic[row][i]) {
			col = i + 1
			break
		}
		if i == 0 {
			col = 0
			break
		}
	}
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

func day3_2(scanner *bufio.Scanner) string {
	//pre declare output at 0
	grandTotal := 0
	//Scanlines has some trouble with the default buffer.
	const maxCapacity int = 19750 // required line length (actually the whole file)
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
		grandTotal += processLineDay3_2(row, line)
	}

	return fmt.Sprint(grandTotal)
}

func processLineDay3_2(row int, line []byte) (rowTotal int) {
	log(string(line))
	for col, char := range line {
		//is it a gear
		if char == '*' {
			gearTotal := 0
			log("gear at ", row, col)
			//find all ajacent part numbers. Use a map so they're unique
			parts := make(map[string]struct{})
			//above
			parts[getPartNumberAt(row-1, col-1)] = struct{}{}
			parts[getPartNumberAt(row-1, col)] = struct{}{}
			parts[getPartNumberAt(row-1, col+1)] = struct{}{}
			//beside
			parts[getPartNumberAt(row, col-1)] = struct{}{}
			parts[getPartNumberAt(row, col+1)] = struct{}{}
			//below
			parts[getPartNumberAt(row+1, col-1)] = struct{}{}
			parts[getPartNumberAt(row+1, col)] = struct{}{}
			parts[getPartNumberAt(row+1, col+1)] = struct{}{}
			//we don't care about non parts
			delete(parts, "")
			log("gear has ", len(parts), " parts")
			//if we have exactly two parts, multiply
			if len(parts) == 2 {
				gearTotal = 1
				for p := range parts {
					val, err := strconv.Atoi(p)
					check(err)
					gearTotal = gearTotal * val
				}

			}
			rowTotal += gearTotal
		}
	}
	log("adding to total: ", rowTotal)
	return
}
