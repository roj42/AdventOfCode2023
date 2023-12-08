package main

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
)

type connectionMap map[byte]string

func day8(scanner *bufio.Scanner) string {

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
	for scanner.Scan() {
		nodeName, nodeItself := parseLineDay8(scanner.Text())
		desertMap[nodeName] = nodeItself

	}

	//Make sure we're not here due to scanner errors
	if err := scanner.Err(); err != nil {
		check(err)
	}

	//walk that route. Count your steps
	stepCount := 0
	repeats := 0
	curNode := desertMap["AAA"]
	for ; ; stepCount++ {
		//do we need to repeat?
		if stepCount == len(route) {
			repeats++
			stepCount = 0
		}
		nextStep := route[stepCount]
		//if we're a step away from the end, hurray
		if curNode[nextStep] == "ZZZ" {
			stepCount++
			break
		}
		curNode = desertMap[curNode[nextStep]]
		log("curnode", curNode)
	}

	//We've arrived! how long did that take?
	grandTotal := repeats*len(route) + stepCount

	return fmt.Sprint("total steps", grandTotal)
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
