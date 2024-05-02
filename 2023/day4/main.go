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

	id := 1
	cardRun := make(map[int]int)
	totalCards := 0

	for scanner.Scan() {
		line := scanner.Text()
		card := strings.Split(line, ":")[1]
		winningNumbersStr := strings.Split(card, "|")[0]
		cardNumbersStr := strings.Split(card, "|")[1]
		cardRun[id]++

		winningNumbers := make(map[string]bool)
		for _, number := range strings.Split(winningNumbersStr, " ") {
			if number != "" {
				winningNumbers[number] = true
			}
		}

		//fmt.Println(id, cardRun[id])

		for range cardRun[id] {
			matches := 0

			for _, number := range strings.Split(cardNumbersStr, " ") {
				if winningNumbers[number] {
					matches++
					cardRun[id+matches]++
					//fmt.Println(id, "| Found match", number, "adding to", id+matches)
				}

			}
		}

		id++
	}

	for _, v := range cardRun {
		totalCards += v
	}
	fmt.Println("Total cards:", totalCards)

}
