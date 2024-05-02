package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	totalPoints := 0
	for scanner.Scan() {
		line := scanner.Text()
		card := strings.Split(line, ":")[1]
		winningNumbersStr := strings.Split(card, "|")[0]
		cardNumbersStr := strings.Split(card, "|")[1]
		points := 0

		winningNumbers := make(map[string]bool)
		for _, number := range strings.Split(winningNumbersStr, " ") {
			if number != "" {
				winningNumbers[number] = true
			}
		}

		for _, number := range strings.Split(cardNumbersStr, " ") {
			if winningNumbers[number] {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}

		}

		totalPoints += points
	}

	fmt.Println("Total points:", totalPoints)
}
