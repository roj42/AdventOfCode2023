package main

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
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

	for _, nodeName := range theAList {
		curNodes = append(curNodes, desertMap[nodeName])

	}

	sparkle := "!@#$%^&*(."
	answers := []int{}
	for i := range curNodes {

		//walk that route. Count your steps
		stepCount := 0
		repeats := 0

		for ; ; stepCount++ {
			//do we need to repeat?
			if stepCount == len(route) {
				repeats++
				stepCount = 0
			}
			nextStep := route[stepCount]
			if (!isPart2 && curNodes[i][nextStep] == "ZZZ") ||
				(isPart2 && curNodes[i][nextStep][2] == 'Z') { // [][][] lol
				fmt.Println("(", string(sparkle[i]), fmt.Sprint(stepCount), "|", fmt.Sprint(repeats), "^", fmt.Sprint((repeats*len(route))+stepCount+1))
				answers = append(answers, (repeats*len(route))+stepCount+1)
				break
			}
			//all together now, step!
			curNodes[i] = desertMap[curNodes[i][nextStep]]
		}
	}
	log("done?")
	grandTotal := LCM(answers[0], answers[1])
	if len(answers) > 2 {
		grandTotal = LCM(answers[0], answers[1], answers[2:]...)
	}

	return fmt.Sprint(grandTotal)
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

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
