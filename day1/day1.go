package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// NOTE: File operation you use the os package
	// the os.Open() function returns a reader for the file
	file, err := os.Open("day1_input.txt")
	if err != nil {
		// NOTE: For loggins and stopping on error use log library
		log.Fatalf("error opening file")
	}
	//NOTE: defer functions are called after surround code has been called
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatalf("error closing file")
		}
	}()

	// NOTE: bufio provides buffer io functions, a low of nice helpers i think
	reader := bufio.NewReader(file)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("error reading line")
	}
	res := make([]string, 0)
	validNumbers := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"9":     "9",
		"8":     "8",
		"7":     "7",
		"6":     "6",
		"5":     "5",
		"4":     "4",
		"3":     "3",
		"2":     "2",
		"1":     "1",
		"0":     "0",
	}
	for err == nil {
		lineRes := ""
		for i := range line {
			for k := range validNumbers {
				if strings.HasPrefix(line[i:], k) {
					lineRes += validNumbers[k]
					break
				}
			}
			if len(lineRes) > 0 {
				break
			}
		}
		for i := range line {
			for k := range validNumbers {
				if strings.HasSuffix(line[:len(line)-1-i], k) {
					lineRes += validNumbers[k]
					break
				}
			}
			if len(lineRes) > 1 {
				break
			}
		}
		line, err = reader.ReadString('\n')
		res = append(res, lineRes)
	}
	finalRes := 0
	for _, s := range res {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("error converting to number")
		}
		finalRes += n
	}
	fmt.Printf("The final answer is %v", finalRes)
}
