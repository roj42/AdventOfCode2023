package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day1(file *os.File) string {
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K?
	lineTotal := 0
	for scanner.Scan() {
		log("line:" + scanner.Text())
		lineTotal += processLineDay1(scanner.Text())
	}

	log("File total: " + fmt.Sprint(lineTotal))

	if err := scanner.Err(); err != nil {
		check(err)
	}
	return fmt.Sprint(lineTotal)
}

func processLineDay1(input string) int {
	first := ""
	last := ""
	log("==LINE " + input + "==")
	msg := "found: "
	for _, char := range input {
		stringchar := string(char)
		_, e := strconv.Atoi(stringchar) //gross
		if e == nil {                    // it's an int
			msg += " " + stringchar
			if first == "" {
				first = stringchar
			}
			last = stringchar
		}
	}
	log(msg)

	total, err := strconv.Atoi(first + last)
	check(err)
	log("==LINE TOTAL " + fmt.Sprint(total) + "==")
	return total

}

func day1_2(file *os.File) string {
	scanner := bufio.NewScanner(file)

	lineTotal := 0
	for scanner.Scan() {
		log("line:" + scanner.Text())

		lineTotal += processLineDay1_2(scanner.Text())
	}

	log("File total: " + fmt.Sprint(lineTotal))

	if err := scanner.Err(); err != nil {
		check(err)
	}
	return fmt.Sprint(lineTotal)
}

func processLineDay1_2(input string) int {
	first := ""
	last := ""
	log("==LINE " + input + "==")
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
	log("==LINE TOTAL " + fmt.Sprint(total) + "==")
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
