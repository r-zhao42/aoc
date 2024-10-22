package main

import (
	"bufio"
	"fmt"
	"log"
	//"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("pog")
	file, err := os.Open("day2_input.txt")
	if err != nil {
		log.Fatalf("error opening file")
	}

	reader := bufio.NewReader(file)
	res := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}
		line = line[:len(line)-1]
		gameAndResults := strings.Split(line, ": ")
		//gameNumberString := strings.Split(gameAndResults[0], " ")
		fmt.Print(line)
		//gameNumber, err := strconv.Atoi(gameNumberString[1])
		//fmt.Println(gameNumber)
		if err != nil {
			log.Fatalf("error converting game number to int")
		}

		gamesResults := strings.Split(gameAndResults[1], "; ")
		maxRed, maxGreen, maxBlue := 0, 0, 0
		fmt.Print(gamesResults)
		for _, game := range gamesResults {
			var cubeNumbers []string
			if strings.Contains(game, ", ") {
				cubeNumbers = strings.Split(game, ", ")
			} else {
				cubeNumbers = []string{game}
			}
			for _, cubeNumber := range cubeNumbers {
				cubeAndNumber := strings.Split(cubeNumber, " ")
				//fmt.Println(cubeAndNumber)
				numCubes, err := strconv.Atoi(cubeAndNumber[0])
				if err != nil {
					log.Fatalf("error converting cube number to int")
				}

				//	fmt.Println(cubeAndNumber)
				switch color := cubeAndNumber[1]; color {
				case "red":
					maxRed = max(maxRed, numCubes)
				case "blue":
					maxBlue = max(maxBlue, numCubes)
				case "green":
					maxGreen = max(maxGreen, numCubes)
				}
				//	fmt.Print(cubeAndNumber[1], numCubes)
			}
			// TODO: Calculate power sum
		}
		powerSet := maxRed * maxBlue * maxGreen
		res += powerSet
	}
	fmt.Println(res)

}
