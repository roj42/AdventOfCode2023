package main

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
	"sync"
)

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
		ghostChans[i] = make(chan int)

	}

	var wg sync.WaitGroup
	//THE LISTENER
	wg.Add(1)
	go func() {
		defer wg.Done()
		var state = make([]int, len(ghostChans))
		highest := -1
		//read one set, record highest
		for i := range state {
			state[i] = <-ghostChans[i]
			if state[i] > highest {
				highest = state[i]
			}
		}

		//start loopin'. If we're equal, great. if not re-fetch all but the highest
		for c := 0; ; c++ {
			if allValuesEqual(state) {
				log("ANSWER", state[0])
				return
			}

			if c%10000 == 0 {
				fmt.Print("()O()", len(ghostChans[0]))
			}
			highNew := -1
			for i := range state {
				if state[i] < highest {
					state[i] = <-ghostChans[i]
					if state[i] > highNew {
						highNew = state[i]

					}
				}
			}
			highest = highNew
		}
	}()

	sparkle := "!@#$%^&*(."
	for i := range curNodes {
		//avoid memory collision here
		i := i
		go func() {
			//walk that route. Count your steps
			stepCount := 0
			repeats := 0
			endCount := 0

			for ; ; stepCount++ {
				//do we need to repeat?
				if stepCount == len(route) {
					repeats++
					stepCount = 0
				}
				nextStep := route[stepCount]
				if (!isPart2 && curNodes[i][nextStep] == "ZZZ") ||
					(isPart2 && curNodes[i][nextStep][2] == 'Z') { // [][][] lol
					ghostChans[i] <- repeats*len(route) + stepCount + 1 //+1 'cause the next step is actually z
					endCount++
					fmt.Print("|", string(stepCount))
					if endCount%10000 == 0 {
						fmt.Print(string(sparkle[i]))
					}
				}
				//all together now, step!
				curNodes[i] = desertMap[curNodes[i][nextStep]]
			}
		}()
	}
	wg.Wait()
	return "\ndone?"
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
