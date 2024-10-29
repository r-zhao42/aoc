package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {

	file, err := os.Open("day3_input.txt")
	if err != nil {
		log.Fatalf("Failed to open file")
	}

	defer file.Close()
	notDigitOrPeriod := regexp.MustCompile("[^0-9.]+")
	digit := regexp.MustCompile("[0-9]+")

	scanner := bufio.NewScanner(file)
	row := 0
	symbolLocations := []Coordinate{}
	digitLocations := make(map[Coordinate]int)
	gearLocations := []Coordinate{}
	for scanner.Scan() {
		line := scanner.Bytes()

		symbolCols := notDigitOrPeriod.FindAllIndex(line, -1)
		for _, location := range symbolCols {
			symbolLocations = append(symbolLocations, Coordinate{row, location[0]})
			if string(line[location[0]]) == "*" {
				gearLocations = append(gearLocations, Coordinate{row, location[0]})
			}
		}

		digitCols := digit.FindAllIndex(line, -1)

		for _, location := range digitCols {
			start, end := location[0], location[1]
			number, err := strconv.Atoi(string(line[start:end]))
			if err != nil {
				log.Fatalf("error converting string to number")
			}
			for i := start; i < end; i++ {
				digitLocations[Coordinate{row, i}] = number
			}
		}
		row += 1
	}
	res := 0
	res2 := 0
	for _, loc := range symbolLocations {
		partNumber := getPartNumber(loc, digitLocations)

		if err == nil {
			res += partNumber
		}
	}

	for _, loc := range gearLocations {
		partNumber := getGearRatio(loc, digitLocations)

		if err == nil {
			res2 += partNumber
		}
	}

	fmt.Println("The part 1 results are ", res)
	fmt.Println("The part 2 results are ", res2)
}

func getPartNumber(loc Coordinate, digitLocs map[Coordinate]int) int {
	// NOTE: Solution only works if adjacent part numbers are unique
	res := 0
	seen := make(map[int]bool)
	for r := loc.Row - 1; r <= loc.Row+1; r++ {
		for c := loc.Col - 1; c <= loc.Col+1; c++ {
			partNum, ok := digitLocs[Coordinate{r, c}]
			_, hasSeen := seen[partNum]
			if ok && !hasSeen {
				res += partNum
				seen[partNum] = true
			}
		}
	}
	return res
}
func getGearRatio(loc Coordinate, digitLocs map[Coordinate]int) int {
	// NOTE: Solution only works if adjacent part numbers are unique
	res := 1
	seen := make(map[int]bool)
	for r := loc.Row - 1; r <= loc.Row+1; r++ {
		for c := loc.Col - 1; c <= loc.Col+1; c++ {
			partNum, ok := digitLocs[Coordinate{r, c}]
			_, hasSeen := seen[partNum]
			if ok && !hasSeen {
				res *= partNum
				seen[partNum] = true
			}
		}
	}
	if len(seen) == 2 {
		return res
	} else {
		return 0
	}
}

type Coordinate struct {
	Row int
	Col int
}
