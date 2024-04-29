package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
)

func isGamePossible(game string) bool {
	records := strings.Split(game, ";")
	for _, record := range records {
		colors := strings.Split(strings.ReplaceAll(record, " ", ""), ",")
		for _, color := range colors {
			if strings.Contains(color, "red") {
				num, _ := strconv.Atoi(strings.ReplaceAll(color, "red", ""))
				if num > maxRed {
					return false
				}
			} else if strings.Contains(color, "green") {
				num, _ := strconv.Atoi(strings.ReplaceAll(color, "green", ""))
				if num > maxGreen {
					return false
				}
			} else if strings.Contains(color, "blue") {
				num, _ := strconv.Atoi(strings.ReplaceAll(color, "blue", ""))
				if num > maxBlue {
					return false
				}
			}
		}
	}

	return true
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	id := 1
	idSum := 0
	for scanner.Scan() {
		line := scanner.Text()
		game := line[strings.IndexRune(line, ':')+2:]

		if isGamePossible(game) {
			idSum += id
		}

		id++
	}

	fmt.Printf("Value: %d\n", idSum)
}
