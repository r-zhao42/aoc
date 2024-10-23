package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func fileReader(fileName string) *bufio.Reader {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening file ")
	}

	reader := bufio.NewReader(file)
	return reader
}

func Slicer(inputSlice *[][]rune, reader *bufio.Reader) map[rune]bool {
	symbols := make(map[rune]bool, 0)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		lineArray := make([]rune, 0)

		for _, c := range line {
			lineArray = append(lineArray, c)
			if !unicode.IsNumber(c) && c != '.' {
				symbols[c] = true
			}

		}
		*inputSlice = append(*inputSlice, lineArray)
	}
	return symbols
}

func getNum(line []rune, col int) (int, int, int) {
	start, end := col, col
	numString := ""
	for unicode.IsNumber(line[end]) && end < len(line) {
		numString += string(line[end])
		end += 1
	}
	num, err := strconv.Atoi(numString)
	if err != nil {
		log.Fatalf("error converting to Number")
	}
	return num, start, end
}

func isPartNum(grid [][]rune, row, start, end int, symbols map[rune]bool) bool {
	prevRowIndex := max(0, row-1)
	nextRowIndex := min(len(grid)-1, row+1)
	prevColIndex := max(0, start-1)
	nextColIndex := min(len(grid[row])-1, end+1)
	// fmt.Println()
	r := prevRowIndex
	for r <= nextRowIndex {
		//for _, c := range grid[r][prevColIndex:nextColIndex] {
		//	// fmt.Print(string(c))
		//}
		// fmt.Println()
		r += 1
	}
	// previous row
	for _, c := range grid[prevRowIndex][prevColIndex:nextColIndex] {
		_, ok := symbols[c]
		if ok {
			return true
		}
	}
	for _, c := range grid[nextRowIndex][prevColIndex:nextColIndex] {
		_, ok := symbols[c]
		if ok {
			return true
		}
	}
	_, ok1 := symbols[grid[row][prevColIndex]]
	_, ok2 := symbols[grid[row][nextColIndex]]
	if ok1 && ok2 {
		return true
	}
	return false
}

func main() {
	reader := fileReader("day3_input.txt")
	inputSlice := make([][]rune, 0)
	symbols := Slicer(&inputSlice, reader)
	row, col := 0, 0
	res := 0
	for c, _ := range symbols {
		fmt.Print(string(c))
	}
	for row < len(inputSlice) {
		line := inputSlice[row]
		col = 0
		for col < len(line) {
			char := inputSlice[row][col]
			// check if char is number, then get the full number, then check if it is a part number
			if unicode.IsNumber(char) {
				// find the number and get the col start and col end indexes of the number
				num, start, end := getNum(line, col)
				isPartNum := isPartNum(inputSlice, row, start, end, symbols)
				// // fmt.Println("number:", num, "isPartNum:", isPartNum)
				if isPartNum {
					res += num
				}
				col = end - 1
			}
			col += 1
		}
		row += 1
	}
	fmt.Println("the answer is :", res)

	// fmt.Println(len(inputSlice))
}
