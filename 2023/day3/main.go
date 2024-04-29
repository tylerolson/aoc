package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"unicode"
)

func isNumberAtIndex(numbers [][]int, i int, j int) int {
	if i >= 0 && i < len(numbers) {
		if j >= 0 && j < len(numbers[i]) {
			if numbers[i][j] != 0 {
				num := numbers[i][j]
				numbers[i][j] = 0

				for k, value := range numbers[i] {
					if value == num {
						numbers[i][k] = 0
					}
				}

				return num
			}
		}
	}

	return -1
}

func checkSymbols(engine [][]rune, numbers [][]int) []int {
	var validNumbers []int
	for i, line := range engine {
		for j, letter := range line {
			if !unicode.IsDigit(letter) && string(letter) != "." {
				if value := isNumberAtIndex(numbers, i-1, j-1); value != -1 {
					validNumbers = append(validNumbers, value)
				}

				if value := isNumberAtIndex(numbers, i, j-1); value != -1 {
					validNumbers = append(validNumbers, value)
				}

				if value := isNumberAtIndex(numbers, i+1, j-1); value != -1 {
					validNumbers = append(validNumbers, value)
				}

				if value := isNumberAtIndex(numbers, i-1, j); value != -1 {
					validNumbers = append(validNumbers, value)
				}

				if value := isNumberAtIndex(numbers, i+1, j); value != -1 {
					validNumbers = append(validNumbers, value)
				}

				if value := isNumberAtIndex(numbers, i-1, j+1); value != -1 {
					validNumbers = append(validNumbers, value)
				}

				if value := isNumberAtIndex(numbers, i, j+1); value != -1 {
					validNumbers = append(validNumbers, value)
				}

				if value := isNumberAtIndex(numbers, i+1, j+1); value != -1 {
					validNumbers = append(validNumbers, value)
				}
			}
		}
	}

	return validNumbers
}

func getNumbers(engine [][]rune) [][]int {
	numbers := make([][]int, len(engine))
	for i := range numbers {
		numbers[i] = make([]int, len(engine[i]))
	}
	var currentNumber string
	var js []int
	for i, line := range engine {
		for j, letter := range line {
			if unicode.IsDigit(letter) {
				currentNumber += string(letter)
				js = append(js, j)
			} else {
				if currentNumber != "" {
					num, _ := strconv.Atoi(currentNumber)
					for _, _j := range js {
						numbers[i][_j] = num
					}

					js = nil
					currentNumber = ""
				}

			}
		}
	}

	return numbers
}

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

	numbers := getNumbers(engine)

	validNumbers := checkSymbols(engine, numbers)

	for _, line := range engine {
		for _, letter := range line {
			fmt.Printf("%c", letter)
		}
		fmt.Println()
	}

	for _, line := range numbers {
		for _, number := range line {
			fmt.Printf("%d ", number)
		}
		fmt.Println()
	}

	sort.Ints(validNumbers)

	fmt.Println(validNumbers)

	finalNumber := 0

	for _, num := range validNumbers {
		finalNumber += num
	}

	fmt.Println(finalNumber)

}
