package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type box []lensLabel

func (b *box) LensMove(newLens lensLabel, op lensOp) {

	for i := range *b {
		if newLens.label == (*b)[i].label {
			if op == ADD {
				(*b)[i].fl = newLens.fl
				return
			} else { //remove
				(*b) = append((*b)[:i], (*b)[i+1:]...)
				return
			}
		}
	}
	//unfound. Add if it's new, otherwise noop
	if op == ADD {
		(*b) = append((*b), newLens)
	}

}

type lensOp byte

const ADD lensOp = '='
const DEL lensOp = '-'

type lensLabel struct {
	label string
	fl    int //focal length
}

func day15(scanner *bufio.Scanner, isPart2 bool) string {

	//day 15 is one giant string
	steps := []string{}
	for scanner.Scan() {
		steps = strings.Split(scanner.Text(), ",")
	}
	//did we error in there somewhere?
	if err := scanner.Err(); err != nil {
		check(err)
	}

	//answer is the sum of hashes
	grandTotal := 0

	if isPart2 {
		var boxArray [256]box
		log("it sure is part 2")
		for _, step := range steps {
			boxNum, lens, op := step2Parse(step)
			// log(boxNum, lens, op)
			boxArray[boxNum].LensMove(lens, op)
		}
		//calculate POWER
		for boxNum, box := range boxArray {
			for slotNum, lens := range box {
				//power is box num (plus 1) times slot num (plus 1) times focal length
				grandTotal += (boxNum + 1) * (slotNum + 1) * lens.fl
			}
		}
	} else {
		for _, step := range steps {
			grandTotal += xmasHash(step)
		}
	}

	return fmt.Sprint(grandTotal)
}

func step2Parse(input string) (boxNum int, lens lensLabel, op lensOp) {
	//split input into parts; example: rn=1, cm-
	//work from the back. last char is always number or nothing, second to last is always the op

	if strings.HasSuffix(input, string(DEL)) {
		op = DEL
		lens.label = input[:len(input)-1]
	} else {
		op = ADD
		lens.label = input[:len(input)-2]
		lens.fl, _ = strconv.Atoi(string(input[len(input)-1]))
	}
	boxNum = xmasHash(lens.label)
	return
}

func xmasHash(step string) int {
	hash := 0
	for _, char := range step {
		hash += int(char)
		hash = (hash * 17) % 256
	}
	// log("hash of", step, "is", hash)
	return hash
}
