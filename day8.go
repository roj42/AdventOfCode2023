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

	//walk that route. Count your steps
	stepCount := 0
	repeats := 0

	//let's get cute with go routines!
	var wg sync.WaitGroup

	//make a list of current nodes based on number of starters
	curNodes := []connectionMap{}
	for _, nodeName := range theAList {
		curNodes = append(curNodes, desertMap[nodeName])

	}

	for ; ; stepCount++ {
		//do we need to repeat?
		if stepCount == len(route) {
			repeats++
			if repeats%1000 == 0 {
				fmt.Print(".")
			}
			stepCount = 0
			// fmt.Print("\nRepeat", repeats, ":")
		}
		nextStep := route[stepCount]
		success := true
		// fmt.Print("\nStep", stepCount, ":")
		var mu sync.Mutex
		for i := range curNodes {
			wg.Add(1)
			//avoid memory collision here
			i := i
			go func() {
				defer wg.Done()
				//race condition against success, but who cares? if anyone fails, we all fail
				if (!isPart2 && curNodes[i][nextStep] != "ZZZ") ||
					(isPart2 && curNodes[i][nextStep][2] != 'Z') { // [][][] lol
					mu.Lock()
					success = false
					mu.Unlock()
				}
				//all together now, step!
				curNodes[i] = desertMap[curNodes[i][nextStep]]

			}()
		}
		wg.Wait()

		if success {
			//it's the NEXT step that succedes
			stepCount++
			break
		}

	}

	//We've arrived! how long did that take?
	grandTotal := repeats*len(route) + stepCount

	return fmt.Sprint("\ntotal steps ", grandTotal)
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
