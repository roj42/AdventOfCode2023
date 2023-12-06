package main

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type gardeningException struct {
	dest   int64
	source int64
	span   int64
}

// create a new exception from a well-ordered list of exactly 3 int64s
func NewGardeningException(input []int64) gardeningException {
	//go will tell us if this didn't work with a handy panic
	return gardeningException{
		dest:   input[0],
		source: input[1],
		span:   input[2],
	}
}

// a gardening map is a list of exceptions
type gardeningMap []gardeningException

func (gm gardeningMap) find(src int64) int64 {
	//first, see if our src is within any span range
	for _, except := range gm {
		//is it greater equal source, but less than the sum of the start+span, it's in!
		if src >= except.source && src < except.source+except.span {
			//hurray, calculate the map. It's the distants (difference) between the span start and our target
			return except.dest + (src - except.source)
		}
	}
	//nope! straight across
	return src
}

func day5(scanner *bufio.Scanner, part2 bool) string {

	//day 5 is different. We'll need a list of seeds, and 7(!) maps
	seeds := []int64{}
	seed2soil := gardeningMap{}
	soil2fert := gardeningMap{}
	fert2agua := gardeningMap{}
	agua2suns := gardeningMap{}
	suns2temp := gardeningMap{}
	temp2damp := gardeningMap{}
	damp2spot := gardeningMap{}
	//Scan off seeds
	if scanner.Scan() {
		log("First Line" + scanner.Text())
		seedsLine := scanner.Text()
		if err := scanner.Err(); err != nil || !strings.HasPrefix(seedsLine, "seeds: ") {
			check(errors.New("first line error. Where are my seeds?"))
		}
		seeds = toBigInts(strings.TrimPrefix(seedsLine, "seeds:"))
	}
	log("seeds: ", seeds)

	//let's load them maps
	//a pointer? why not. Pointers are cool.
	var curMap *gardeningMap
	for scanner.Scan() {
		curLine := scanner.Text()
		if curLine == "" {
			continue
		}

		log("line: ", curLine)
		if !isDigit(curLine[0]) {
			//new map! Exciting
			mapTitle := strings.Split(curLine, ":")[0]
			switch mapTitle {
			case "seed-to-soil map":
				curMap = &seed2soil
			case "soil-to-fertilizer map":
				curMap = &soil2fert
			case "fertilizer-to-water map":
				curMap = &fert2agua
			case "water-to-light map":
				curMap = &agua2suns
			case "light-to-temperature map":
				curMap = &suns2temp
			case "temperature-to-humidity map":
				curMap = &temp2damp
			case "humidity-to-location map":
				curMap = &damp2spot
			}
			continue
		}
		//capture the conversion exception
		*curMap = append(*curMap, NewGardeningException(toBigInts(curLine)))
	}

	//scanner is weird about errors. It will kick us out of the loop that .Scan() produces when there is one, so do this odd check.
	if err := scanner.Err(); err != nil {
		check(err)
	}

	//Alright. Let's map them seeds
	lowestResult := int64(9223372036854775807)
	if part2 {
		log("+ is a start of a new seed range (of 10), dots are 1000000 seeds")
		for i, seedPart := range seeds {

			//part 2 has MOAR SEEDS
			//skip evens, starting at 0
			if i%2 == 0 {
				continue
			}
			fmt.Print("\n+")
			seedStart := seeds[i-1]
			seedEnd := seedStart + seedPart
			for seed := seedStart; seed < seedEnd; seed++ {
				if seed%1000000 == 0 {
					fmt.Print(".")
				}
				//I could get one-line-cute, here, but the compiler will do it for me. Thanks, compiler!
				soil := seed2soil.find(seed)
				fert := soil2fert.find(soil)
				agua := fert2agua.find(fert)
				suns := agua2suns.find(agua)
				temp := suns2temp.find(suns)
				damp := temp2damp.find(temp)
				spot := damp2spot.find(damp)
				if spot < lowestResult {
					lowestResult = spot
				}
			}
		}
		fmt.Print("\n")
	} else {
		for _, seed := range seeds {

			//I could get one-line-cute, here, but the compiler will do it for me. Thanks, compiler!
			soil := seed2soil.find(seed)
			fert := soil2fert.find(soil)
			agua := fert2agua.find(fert)
			suns := agua2suns.find(agua)
			temp := suns2temp.find(suns)
			damp := temp2damp.find(temp)
			spot := damp2spot.find(damp)
			if spot < lowestResult {
				lowestResult = spot
			}
		}
	}

	return fmt.Sprint(lowestResult)
}

func toBigInts(stringOfInts string) []int64 {
	listOfInts := []int64{}
	for _, str := range strings.Split(stringOfInts, " ") {
		//note that we might have bonus spaces for single digit numbers, so skip those
		if str == "" {
			continue
		}
		parsedInt, e := strconv.ParseInt(str, 10, 64)
		check(e)
		listOfInts = append(listOfInts, parsedInt)
	}
	return listOfInts
}
