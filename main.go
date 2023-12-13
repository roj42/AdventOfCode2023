package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
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
		testPrefix = os.Args[2]
	}

	start := time.Now()

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
	case "6":
		fmt.Println(day6(scanner, false))
	case "6_2":
		fmt.Println(day6(scanner, true))
	case "7":
		fmt.Println(day7(scanner, false))
	case "7_2":
		fmt.Println(day7(scanner, true))
	case "8":
		fmt.Println(day8(scanner, false))
	case "8_2":
		fmt.Println(day8(scanner, true))
	case "9":
		fmt.Println(day9(scanner, false))
	case "9_2":
		fmt.Println(day9(scanner, true))
	case "10":
		fmt.Println(day10(scanner, false))
	case "10_2":
		fmt.Println(day10(scanner, true))
	case "11":
		fmt.Println(day11(scanner, false))
	case "11_2":
		fmt.Println(day11(scanner, true))
	default:
		log("no implementation for day: " + dayInput)

	}
	stop := time.Since(start)
	log("time", stop.String())

}
