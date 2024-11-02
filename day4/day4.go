package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error opening the file")
	}

	scanner := bufio.NewScanner(file)

	res := 0
	lineNumber := 0
	cardMultiplierSlice := make([]int, 1)
	for scanner.Scan() {
		line := scanner.Text()

		card := strings.Split(line, ": ")[1]

		cardResults := strings.Split(card, " | ")

		digitRegex, err := regexp.Compile("\\d+")
		if err != nil {
			log.Fatal("Error compiling regex")
		}

		winningNumbers := digitRegex.FindAll([]byte(cardResults[0]), -1)
		myNumbers := digitRegex.FindAll([]byte(cardResults[1]), -1)

		winningHash := make(map[int]bool)
		for _, n := range winningNumbers {
			num, err := strconv.Atoi(string(n))
			if err != nil {
				fmt.Println(n)
				log.Fatal("Error converting to number")
			}
			winningHash[num] = true
		}

		matches := 0
		for _, n := range myNumbers {
			num, err := strconv.Atoi(string(n))
			if err != nil {
				fmt.Println(n)
				log.Fatal("Error converting to number")
			}

			_, ok := winningHash[num]
			if ok {
				matches++
			}
		}
		// NOTE: part 1 solution
		// res += part1(matches)

		// NOTE: part 2 solutions
		if lineNumber >= len(cardMultiplierSlice) {
			cardMultiplierSlice = append(cardMultiplierSlice, 1)
		} else {
			cardMultiplierSlice[lineNumber] += 1
		}
		for i := lineNumber + 1; i <= lineNumber+matches; i++ {
			if i >= len(cardMultiplierSlice) {
				cardMultiplierSlice = append(cardMultiplierSlice, 1)
			} else {
				cardMultiplierSlice[i] += cardMultiplierSlice[lineNumber]
			}
		}
		res += cardMultiplierSlice[lineNumber]

		lineNumber++
	}

	fmt.Println("Part 2 result:", res)
}
func part1(matches int) int {
	points := 0
	if matches != 0 {
		points = int(math.Pow(2, float64(matches-1)))
	}

	return points
}
