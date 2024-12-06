package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputText, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal("Error reading file")
	}

	inputs := strings.Split(string(inputText), "\n\n")

	seeds := make([]int, 0)
	maps := make([]map[int][]int, 0)

	for _, mapText := range inputs {
		mapArray := strings.Split(mapText, ":")

		mapName, mapContent := mapArray[0], mapArray[1]

		if mapName == "seeds" {
			seeds = processSeeds(mapContent)
		} else {
			maps = append(maps, processMaps(mapContent))
		}
	}

	partOneRes := findMinLocation(seeds, maps)

	fmt.Println("Part 1 solution:", partOneRes)

}

func findMinLocation(seeds []int, maps []map[int][]int) int {
	minRes := math.MaxInt64
	for _, seed := range seeds {
		seedVal := seed
		for _, currMap := range maps {
			for src, values := range currMap {
				dst, rnge := values[0], values[1]
				if src <= seedVal && seedVal < src+rnge {
					seedVal = dst + (seedVal - src)
					break
				}
			}
		}
		if seedVal < minRes {
			minRes = seedVal
		}
	}

	return minRes
}

func processMaps(mapText string) map[int][]int {
	mapLines := strings.Split(strings.TrimSpace(mapText), "\n")
	res := make(map[int][]int)

	for _, line := range mapLines {
		lineValues := strings.Split(line, " ")

		src, err1 := strconv.Atoi(lineValues[1])
		dest, err2 := strconv.Atoi(lineValues[0])
		rnge, err3 := strconv.Atoi(lineValues[2])

		if err1 != nil || err2 != nil || err3 != nil {
			log.Fatal("error converting map text to number")
		}

		res[src] = []int{dest, rnge}
	}

	return res
}

func processSeeds(seedText string) []int {
	seeds := strings.SplitAfter(strings.Trim(seedText, " "), " ")
	res := make([]int, 0)
	for _, s := range seeds {
		seedNum, err := strconv.Atoi(strings.TrimSpace(s))

		if err != nil {
			log.Fatal("Error converting to number for seed: ", err)
		}
		res = append(res, seedNum)
	}
	return res
}
