package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"strings"
)

const sparkle = "!@#$%^&*("

type connectionMap map[byte]string

func day8(scanner *bufio.Scanner, isPart2 bool) string {

	//Scan off our route
	scanner.Scan()
	route := scanner.Text()
	log("route:", route)

	//skip the blank line
	scanner.Scan()
	if scanner.Text() != "" {
		check(errors.New("we skipped a real line, tiger"))
	}

	//a map that is a literal map!
	desertMap := map[string]connectionMap{}
	theAList := []string{}
	for scanner.Scan() {
		nodeName, nodeItself := parseLineDay8(scanner.Text())
		if isPart2 && nodeName[2] == 'A' {
			theAList = append(theAList, nodeName)
		}
		desertMap[nodeName] = nodeItself

	}

	//Make sure we're not here due to scanner errors
	if err := scanner.Err(); err != nil {
		check(err)
	}

	if !isPart2 && len(theAList) == 0 {
		theAList = append(theAList, "AAA")
	}

	//let's get cute with go routines!

	//make a list of current nodes based on number of starters
	curNodes := []connectionMap{}
	ghostChans := make([]chan int, len(theAList))

	for i, nodeName := range theAList {
		curNodes = append(curNodes, desertMap[nodeName])
		ghostChans[i] = make(chan int, 10000)

	}

	answers := make([]int, len(curNodes))
	for i := range curNodes {

		//New method: record our steps, and then see if we've looped. Assume all must loop.
		stepCount := 0
		repeats := 0
		pathRecord := []byte{}

		//let's find one loop at a time
		for ; ; stepCount++ {
			//was the answer found?
			if answers[i] != 0 {
				fmt.Print("(" + string(sparkle[i]) + "OUT)")
				break
			}
			//do we need to repeat?
			if stepCount == len(route) {
				repeats++
				if repeats%100000 == 0 {
					// fmt.Print(len(pathRecord), string(sparkle[i]))
					fmt.Println(string(pathRecord))
				}
				stepCount = 0
			}
			//peek next step, and record our route
			nextStep := route[stepCount]
			//store the route
			pathRecord = append(pathRecord, curNodes[i][nextStep]...)

			//is that next step a Z?
			if (!isPart2 && curNodes[i][nextStep] == "ZZZ") ||
				(isPart2 && curNodes[i][nextStep][2] == 'Z') { // [][][] lol
				totalSteps := repeats*len(route) + stepCount + 1 //+1 'cause the next step is actually z
				//Now check, if we've taken an even number of steps, if we've looped
				//we're in a loop if the front half of our record exactly matches the back.
				if totalSteps%2 == 0 && bytes.Equal(pathRecord[:len(pathRecord)/2], pathRecord[len(pathRecord)/2:]) {
					answers[i] = totalSteps / 2
				}
			}
			//all together now, step!
			curNodes[i] = desertMap[curNodes[i][nextStep]]
		}

	}
	grandTotal := 1
	for _, ans := range answers {
		grandTotal = grandTotal * ans
	}
	return fmt.Sprint("\n", grandTotal)
}

func parseLineDay8(input string) (nodeName string, connections connectionMap) {
	// example AAA = (BBB, BBB)
	//split at the equals
	equalParts := strings.Split(input, "=")
	//ignore spaces on the left part for the name
	nodeName = strings.Trim(equalParts[0], " ")
	//Ignore spaces and parens to get the two node connections, then split on comma
	nodeParts := strings.Split(equalParts[1], ",")
	connections = make(connectionMap, 2)
	connections['L'] = strings.Trim(nodeParts[0], " ()")
	connections['R'] = strings.Trim(nodeParts[1], " ()")
	return
}

func allValuesEqual(state []int) bool {
	val := state[0]
	for _, v := range state {
		if val != v {
			return false
		}
	}
	return true
}
