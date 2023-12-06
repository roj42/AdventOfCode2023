package main

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func day6(scanner *bufio.Scanner, part2 bool) string {

	times := []int64{}

	//Scan in our data, it's only two lines
	scanner.Scan()
	log("Time Line> " + scanner.Text())
	timeData := scanner.Text()
	if err := scanner.Err(); err != nil || !strings.HasPrefix(timeData, "Time: ") {
		check(errors.New("first line error. Where is time?"))
	}
	if part2 {
		times = append(times, mashint64s(strings.TrimPrefix(timeData, "Time:")))
	} else {
		times = toBigInts(strings.TrimPrefix(timeData, "Time:"))
	}
	log("times: ", times)

	distances := []int64{}
	scanner.Scan()
	log("Dist Line> " + scanner.Text())
	distData := scanner.Text()
	if err := scanner.Err(); err != nil || !strings.HasPrefix(distData, "Distance: ") {
		check(errors.New("first line error. Where is distance?"))
	}
	if part2 {
		distances = append(distances, mashint64s(strings.TrimPrefix(distData, "Distance:")))
	} else {
		distances = toBigInts(strings.TrimPrefix(distData, "Distance:"))
	}

	log("distances: ", distances)
	if len(times) != len(distances) {
		check(errors.New("mismatch times and distances"))
	}
	log(len(times), "times and", len(distances), "distances")

	grandTotal := int64(1)
	//find the number of descreet seconds that can beat the distance record, and multiply those together
	for i := range times {
		grandTotal = grandTotal * numWaysBeatDistForTime(times[i], distances[i])
	}

	return fmt.Sprint(grandTotal)
}

func numWaysBeatDistForTime(time, dist int64) int64 {
	//let me count the ways
	log("A", time, "ms race. Record is", dist, "mm")
	//search from one second to one second less than time, since those will be distance 0
	fmt.Print("Tests from bottom: ")
	bottomLosses := int64(0)
	for i := int64(1); i < time; i++ {
		//calc time
		speed := i                     //in mm/ms
		calcTime := speed * (time - i) //I held it down for i, so I have time-i ms left, during which I go speed
		//if it wins, count it
		if calcTime > dist {
			fmt.Print("+")
			bottomLosses = i
			break
		} else if i%1000000 == 0 {
			fmt.Print(".")
		}
	}
	fmt.Println("\n", bottomLosses, " ways you'll lose")

	fmt.Print("Tests from top: ")
	topLosses := int64(0)
	for i := time; i > bottomLosses; i-- {
		//calc time
		speed := i                     //in mm/ms
		calcTime := speed * (time - i) //I held it down for i, so I have time-i ms left, during which I go speed
		//if it wins, count it
		if calcTime > dist {
			fmt.Print("+")
			topLosses = time - i
			break
		} else if i%1000000 == 0 {
			fmt.Print(".")
		}
	}
	fmt.Println("\n", topLosses, " ways you'll lose")

	//number of possiblities is the time, plus 1 for 0, minus the number of losses on the bottom, and top
	log((time + 1), "possibilities, minus ", topLosses, "and", bottomLosses)
	countWays := (time + 1) - topLosses - bottomLosses
	log(" ", countWays, "!")
	return countWays
}

func mashint64s(input string) int64 {
	mash, e := strconv.ParseInt(strings.Join(strings.Split(input, " "), ""), 10, 64)
	log("mashed", input, "to", mash)
	check(e)
	return mash
}
