package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func log(inputs ...any) {
	fmt.Println(inputs)
}

func main() {

	//did we remember to put an arg of the day?
	if len(os.Args) < 2 || os.Args[1] == "" {
		panic("you forgot the day, sport")
	}
	dayInput := os.Args[1]

	filePrefix := strings.Split(dayInput, "_")[0]
	//open the day's file, and close it when we're done with main, here.
	fileName := "./data/" + filePrefix + ".txt"

	file, err := os.Open(fileName)
	check(err)
	log("opened " + fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, you canresize scanner's capacity for lines over 64K? I hope I never need this note.

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

	default:
		log("no implementation for day: " + dayInput)
	}

}
