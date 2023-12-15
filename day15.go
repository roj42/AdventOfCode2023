package main

import (
	"bufio"
	"fmt"
)

func day15(scanner *bufio.Scanner, isPart2 bool) string {

	grandTotal := 0
	//day 14 is one grid
	platform := platform{}
	for scanner.Scan() {
		//scan until blank line
		platform = append(platform, scanner.Text())
	}
	//did we error in there somewhere?
	if err := scanner.Err(); err != nil {
		check(err)
	}

	if isPart2 {
		log("it sure is part 2")
		//start cyclin
		snapShots := []string{}
		snapShot := ""
		for _, line := range platform {
			snapShot = snapShot + line
		}
		snapShots = append(snapShots, snapShot)

		for i := 0; i < 1000; i++ {

			platform.tilt(UP)
			platform.tilt(LEFT)
			platform.tilt(DOWN)
			platform.tilt(RIGHT)
			snapShot := ""
			for _, line := range platform {
				snapShot = snapShot + line
			}
			if i%10000 == 0 {
				for j, ss := range snapShots {
					if snapShot == ss {
						log("cycle", i, "is the same as", j)
						// break
						return fmt.Sprint(platform.weigh(UP))
					}
				}
			}
			snapShots = append(snapShots, snapShot)
			if len(snapShots) > 1000 {
				snapShots = snapShots[:1000]
			}
		}
	} else {
		platform.tilt(UP)
	}

	grandTotal += platform.weigh(UP)

	return fmt.Sprint(grandTotal)
}
