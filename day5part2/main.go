package main

/*
Input is a text file
First line is seeds that starts with "seeds:" and then has pairs of numbers with first being the starting number and second being the range
e.g. seeds: 79 14 ==> 79, 80, 81, ... 93

Map blocks delimited by new line.
of form:
mapname map:\n
source destination range
e.g.
seed-to-soil map:
50 98 2				=> 50,51 => 98, 99


We want to run each range through all the maps and to find the lowest location.
NOTE: The maps are organized in the order that we want need the pass them through. Potentially able to do this in place



Other input is maps

seeds: 50 18

seed-to-soil map:
55 98 2
59 50 48

In a map range:
55, 56 => 98, 99
59...68 => 50...59
Leftovers:
50...54 => 50...54
57, 58 => 57, 58

NOTE: For each seed at each map, need to keep track of every range that has been mapped and then after comparing to each value in the map, need to create new ranges based on parts of the seed range that haven't been mapped

From this example, we see we need to iterate through all the values in the map

We map the values in the range to the corresponding destination
Given the size of the ranges in the input txt, it is infeasible to map each number in the range one by one.


we will have to do range based calculations

Easier to calculate based on the range and see which numbers are in the seed
e.g.
map: 55 98 2
seeds 50 18

srcStart, srcEnd := map[0], map[0] + map[2]
seedStart, seedEnd := seed[0], seed[0] + seed[1]

if seedEnd < srcStart or seedStart > srcEnd { do nothing }
else {
	return (max(seedStart,srcStart), min(seedEnd, srcEnd))
}
*/

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal("error reading file")
	}

	txtInputs := strings.Split(string(file), "\n\n")

	seedTxt := txtInputs[0]
	mapsTxt := txtInputs[1 : len(txtInputs)-1]

	seeds := processSeeds(seedTxt)
	maps := processMaps(mapsTxt)
	// TODO: Maybe I can simplify the problem by simplifying the ranges somehow?
	for _, m := range maps {
		newRanges := make([][]int, 0)

		for _, s := range seeds {
			seedStart := s[0]
			seedEnd := s[0] + s[1]

			srcRanges := make([][]int, 0)
			for _, r := range m {
				srcStart := r[0]
				srcEnd := r[0] + r[2]

				dstStart := r[1]
				dstEnd := r[1] + r[2]

				if seedEnd < srcStart || seedStart > srcEnd {
					continue
				} else {
					newRange := getRangeForDst(seedStart, seedEnd, srcStart, srcEnd, dstStart, dstEnd)
					newRanges = append(newRanges, newRange)

					srcRange := []int{max(srcStart, seedStart), min(srcEnd, seedEnd)}
					srcRanges = append(srcRanges, srcRange)

				}
			}
			// need to add ranges that didn't intersect with anything in the map as themselves
			/*
				5, 20 ==> (5, 25)

				7 4 ==> (7, 11)
				15 2 ==> (15, 17)



				5 - 7 ==> 5, 2
				11 - 14 ==> 11, 4
				17 - 25 ==> 17, 4

				for any range, if there are n intersecting rances,
			*/
			newSrcRanges := [][]int{[]int{seedStart, seedEnd}}

			newNewSrcRanges := make([][]int, 0)
			for _, sr := range srcRanges {
				srStart := sr[0]
				srEnd := sr[1] + sr[1]

				for _, nsr := range newSrcRanges {
					nr := []int{max()}
					newNewSrcRanges = append(newNewSrcRanges, asdf)
				}

			}
			test = asdf

		}
	}
}

func getRangeForDst(seedStart, seedEnd, srcStart, srcEnd, dstStart, dstEnd int) []int {
	var outStart, outEnd int

	if seedStart > srcStart {
		outStart = dstStart + (seedStart - srcStart)
	} else {
		outStart = dstStart
	}

	if seedEnd < srcEnd {
		outEnd = dstStart + (seedEnd - srcStart)
	} else {
		outEnd = dstEnd
	}

	return []int{outStart, outEnd}
}

func processMaps(mapTextArrays []string) [][][]int {
	res := make([][][]int, 0)

	for _, mapText := range mapTextArrays {
		mapLines := strings.Split(mapText, "\n")[1:]
		currMap := make([][]int, 0)
		for _, line := range mapLines {
			lineSplit := strings.Split(line, " ")
			srcStart, err1 := strconv.Atoi(lineSplit[0])
			dstStart, err2 := strconv.Atoi(lineSplit[1])
			mapRange, err3 := strconv.Atoi(lineSplit[2])

			if err1 != nil || err2 != nil || err3 != nil {
				log.Fatal("Error converting string to number")
			}

			currMap = append(currMap, []int{srcStart, dstStart, mapRange})
		}
		res = append(res, currMap)
	}
	return res
}

func processSeeds(seedTxt string) [][]int {
	seedsNumString := strings.Split(seedTxt, ": ")[1]

	seedsStringArray := strings.Split(seedsNumString, " ")

	seeds := make([][]int, 0)

	for i := 0; i < len(seedsStringArray); i += 2 {
		seedStart, err := strconv.Atoi(seedsStringArray[i])
		if err != nil {
			log.Fatal("Error converting number")
		}

		seedEnd, err := strconv.Atoi(seedsStringArray[i+1])
		if err != nil {
			log.Fatal("Error converting number")
		}

		seeds = append(seeds, []int{seedStart, seedEnd})
	}

	return seeds
}
