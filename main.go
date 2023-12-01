package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func log(s string) {
	fmt.Println(s)
}

func main() {

	//did we remember to put an arg of the day?
	if len(os.Args) < 2 || os.Args[1] == "" {
		panic("you forgot the day, sport")
	}
	dayInput := os.Args[1]

	//open the day's file, and close it when we're done with main, here.
	fileName := "./data/" + dayInput + ".txt"

	file, err := os.Open(fileName)
	check(err)
	log("opened " + fileName)
	defer file.Close()

	//WHAT DAY IS IT?
	switch dayInput {
	case "1":
		fmt.Println(day1(file))
	case "1_2":
		fmt.Println(day1_2(file))

	default:
		log("no implementation for day: " + dayInput)
	}

}
