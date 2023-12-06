package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// proccess day 1. On each line, pull out the "outsidemost" digits, so e14gggt9 is 1 and 9 so 19.
// A single digit is both left and right outermost, natch
func day1(scanner *bufio.Scanner) string {

	//pre declare output at 0
	grandTotal := 0
	//fancy scanner iteration
	for scanner.Scan() {
		log("line:", scanner.Text())
		//process each line in a sub fucntion so I don't go nuts.
		grandTotal += processLineDay1(scanner.Text())
	}

	//scanner is weird about errors. It will kick us out of the loop that .Scan() produces if there is one, so yay
	if err := scanner.Err(); err != nil {
		check(err)
	}

	//just noticed I log and then output it at the end as well. Whatever.
	log("File total: ", grandTotal)
	return fmt.Sprint(grandTotal)
}

// process line will turn a single line into the outermost sum
func processLineDay1(input string) int {
	//placeholder first one
	first := ""
	//placeholder second one
	last := ""
	log("==LINE ", input, "==")
	msg := "found: "
	//range is neat, and produces the for-suff for us. Here it gives us "runes" of each string
	for _, char := range input { //I obviously think of runes as characters
		stringchar := string(char)       //convert back to a string onces, we use this string form a bunch.
		_, e := strconv.Atoi(stringchar) //gross. Try to convert to an int. We don't care what the number is, but if there's no error...
		if e == nil {                    // it's an int
			msg += " " + stringchar //log it
			if first == "" {        // save it if it's leftmost
				first = stringchar
			}
			last = stringchar //this is the rightmost thus far
		}
	} //done parsing the line
	log(msg)

	//so, slap the first and last digits together, and then convert THAT to an integer
	total, err := strconv.Atoi(first + last)
	//man, I hope that worked. Vomit if not.
	check(err)
	log("==LINE TOTAL ", total, "==")
	return total

}

func day1_2(scanner *bufio.Scanner) string {

	grandTotal := 0
	for scanner.Scan() {
		log("line:", scanner.Text())

		grandTotal += processLineDay1_2(scanner.Text())
	}

	log("File total: ", grandTotal)

	if err := scanner.Err(); err != nil {
		check(err)
	}
	return fmt.Sprint(grandTotal)
}

func processLineDay1_2(input string) int {
	first := ""
	last := ""
	log("==LINE ", input, "==")
	msg := "found: "
	for i, char := range input {
		valueToAdd := ""
		stringchar := string(char)
		_, e := strconv.Atoi(stringchar) //gross
		if e == nil {                    // it's an int
			valueToAdd = stringchar

		} else {
			// it's not an int... is it a spelled out word?
			maxlen := i + 5
			if len(input) < maxlen {
				maxlen = len(input)
			}
			valueToAdd = wordToDigit(input[i:maxlen])
			//we're losing some process by checking the later words again, but fuck it
		}
		if valueToAdd != "" {
			msg += " " + valueToAdd
			if first == "" {
				first = valueToAdd
			}
			last = valueToAdd
		}
	}
	log(msg)

	total, err := strconv.Atoi(first + last)
	check(err)
	log("==LINE TOTAL ", fmt.Sprint(total), "==")
	return total

}

// scans a string for the descreet spelling of "one" to "nine" and
// returns the single character numerical equivalent, "1" to "9"
// blank if none are found
func wordToDigit(input string) string {
	if strings.HasPrefix(input, "one") {
		return "1"
	}
	if strings.HasPrefix(input, "two") {
		return "2"
	}
	if strings.HasPrefix(input, "three") {
		return "3"
	}
	if strings.HasPrefix(input, "four") {
		return "4"
	}
	if strings.HasPrefix(input, "five") {
		return "5"
	}
	if strings.HasPrefix(input, "six") {
		return "6"
	}
	if strings.HasPrefix(input, "seven") {
		return "7"
	}
	if strings.HasPrefix(input, "eight") {
		return "8"
	}
	if strings.HasPrefix(input, "nine") {
		return "9"
	}
	return ""
}
