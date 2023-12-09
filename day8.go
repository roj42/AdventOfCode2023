package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"math"
	"strings"
	"sync"
)

const sparkle = "!@#$%^&*("

type connectionMap map[byte]string
type Container struct {
	mu      sync.Mutex
	answers []int
}

func (c *Container) init() {
	for i := range c.answers {
		c.answers[i] = math.MaxInt
	}
}

func (c *Container) submit(i, amt int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if amt < c.answers[i] {
		fmt.Print("(" + string(sparkle[i]) + " is done: " + fmt.Sprint(amt) + ")")
		c.answers[i] = amt
	}
}

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

	var wg sync.WaitGroup
	c := Container{
		answers: make([]int, len(curNodes)),
	}

	c.init()
	for i := range curNodes {

		// wg.Add(1)
		// //avoid memory collision here
		// i := i
		// go func() {
		// defer wg.Done()
		//New method: record our steps, and then see if we've looped
		stepCount := 0
		repeats := 0
		pathRecord := []byte{}

		for ; ; stepCount++ {
			//was the answer found?
			if c.answers[i] != math.MaxInt {
				fmt.Print("(" + string(sparkle[i]) + "OUT)")
				break
			}
			//do we need to repeat?
			if stepCount == len(route) {
				repeats++
				if repeats%100000 == 0 {
					fmt.Print(string(sparkle[i]))
					wg.Wait()
				}
				stepCount = 0
			}
			//peek next step, and record our route
			nextStep := route[stepCount]
			//convert name of the node for storage
			pathRecord = append(pathRecord, curNodes[i][nextStep]...)

			//is that next step a Z?
			if (!isPart2 && curNodes[i][nextStep] == "ZZZ") ||
				(isPart2 && curNodes[i][nextStep][2] == 'Z') { // [][][] lol
				totalSteps := repeats*len(route) + stepCount + 1 //+1 'cause the next step is actually z
				//Now check, if we've taken an even number of steps, if we've looped
				//we're in a loop if the front half of our record exactly matched the back.
				//send a go routine to check this out
				ts := totalSteps
				//let a go routine do the heavy lifting so we can race ahead
				wg.Add(1)
				go func(pr []byte) {
					defer wg.Done()
					if ts%2 == 0 && bytes.Equal(pr[:len(pr)/2], pr[len(pr)/2:]) {
						c.submit(i, ts/2)
					}

					// fmt.Print(".")
				}(pathRecord)
			}
			//all together now, step!
			curNodes[i] = desertMap[curNodes[i][nextStep]]
		}

	}
	wg.Wait()
	grandTotal := 1
	for _, ans := range c.answers {
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
