package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var engine [][]rune
	row := 0
	for scanner.Scan() {
		engine = append(engine, make([]rune, len(scanner.Text()))) //make new empty row size of line

		for i, v := range scanner.Text() {
			engine[row][i] = v
		}
		row++
	}

	for _, line := range engine {
		for _, letter := range line {
			fmt.Printf("%c", letter)
		}
		fmt.Println()
	}

}
