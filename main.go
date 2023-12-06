package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// don't do this in real life
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// don't do this in real life
// actually this is PROBABLY kind of okay.
func log(inputs ...any) {
	fmt.Println(inputs...)
}

func main() {

	//did we remember to put an arg of the day?
	if len(os.Args) < 2 || os.Args[1] == "" {
		panic("you forgot the day, sport")
	}
	dayInput := os.Args[1]

	//any second prefix means load the test file
	testPrefix := ""
	if len(os.Args) > 2 && os.Args[2] != "" {
		testPrefix = "t"
	}

	filePrefix := strings.Split(dayInput, "_")[0]
	//open the day's file, and close it when we're done with main, here.
	fileName := "./data/" + filePrefix + testPrefix + ".txt"

	file, err := os.Open(fileName)
	check(err)
	log("opened " + fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, you can resize scanner's capacity for lines over 64K? I hope I never need this note.

	//WHAT DAY IS IT?
	switch dayInput {
	case "1":
		fmt.Println(day1(scanner))
	case "1_2":
		fmt.Println(day1_2(scanner))
	case "2":
		fmt.Println(day2(scanner))
	case "2_2":
		fmt.Println(day2_2(scanner))
	case "3":
		fmt.Println(day3(scanner))
	case "3_2":
		fmt.Println(day3_2(scanner))
	case "4":
		fmt.Println(day4(scanner))
	case "4_2":
		fmt.Println(day4_2(scanner))
	case "5":
		fmt.Println(day5(scanner, false))
	case "5_2":
		fmt.Println(day5(scanner, true))
	default:
		log("no implementation for day: " + dayInput)
	}

}
