package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type cubeMax struct {
	red   int
	blue  int
	green int
}

var cubeTarget = cubeMax{
	red:   12,
	green: 13,
	blue:  14,
}

func day2(scanner *bufio.Scanner) string {

	//pre declare output at 0
	grandTotal := 0

	for scanner.Scan() {
		log("line:" + scanner.Text())

		//Line step: parse each line for sweet sweet details
		gameNum, lineCubeMax := processLineDay2(scanner.Text())
		log(fmt.Sprint("parsed Line: ", gameNum, lineCubeMax))

		//Aggregation step: do we add? what are we adding?
		//if the line is less than or equal to our max, we can add it.
		if cubeTarget.red >= lineCubeMax.red && cubeTarget.green >= lineCubeMax.green && cubeTarget.blue >= lineCubeMax.blue {
			log(fmt.Sprint("Game ", gameNum, " seems cool"))
			grandTotal += gameNum
		} else {
			log(fmt.Sprint("Game ", gameNum, " NOT cool"))
		}

	}

	//scanner is weird about errors. It will kick us out of the loop that .Scan() produces when there is one, so do this odd check.
	if err := scanner.Err(); err != nil {
		check(err)
	}

	return fmt.Sprint(grandTotal)
}

func day2_2(scanner *bufio.Scanner) string {

	//pre declare output at 0
	grandTotal := 0

	for scanner.Scan() {
		log("line:", scanner.Text())

		//Line step: parse each line for sweet sweet details
		gameNum, lineCubeMax := processLineDay2(scanner.Text())
		log("parsed Line: ", gameNum, lineCubeMax)

		//Aggregation step: do we add? what are we adding?
		//the line maxes are our minimum numbers we'd need. calculate the "power" by multiplying together.
		power := lineCubeMax.red * lineCubeMax.green * lineCubeMax.blue
		log("power: ", power)

		grandTotal += power

	}

	//scanner is weird about errors. It will kick us out of the loop that .Scan() produces when there is one, so do this odd check.
	if err := scanner.Err(); err != nil {
		check(err)
	}

	return fmt.Sprint(grandTotal)
}

// parse the line for the number of the game
// and the maximum you've seen each color, stuffed into result
func processLineDay2(input string) (gameNum int, result cubeMax) {
	//be warned, we're gonna do a lot of input trusting here
	gameAndReveals := strings.Split(input, ":")
	//gameAndReveals of 0 is like "Game 1"

	//split out game number:
	gameAndNum := strings.Split(gameAndReveals[0], " ")
	gameNum, err := strconv.Atoi(gameAndNum[1])
	check(err)

	//Split out reveals
	revealsList := strings.Split(gameAndReveals[1], ";")
	//a reveal is one or more cube pulls
	for _, reveal := range revealsList {
		//a cube pull is like "5 blue"
		cubePulls := strings.Split(reveal, ",")
		for _, pull := range cubePulls {
			numAndColor := strings.Split(pull, " ")
			//there's a leading space, so it'll look like "", "4","red"
			num, e := strconv.Atoi(numAndColor[1])
			check(e)
			color := numAndColor[2]
			switch color {
			case "red":
				if num > result.red {
					result.red = num
				}
			case "green":
				if num > result.green {
					result.green = num
				}
			case "blue":
				if num > result.blue {
					result.blue = num
				}
			}
		}
	}
	return
}

//samples
//Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
// Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
// Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
// Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
// Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
