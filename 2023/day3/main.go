package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func isNumberAtIndex(numbers [][]int, i int, j int) int {
	if i >= 0 && i < len(numbers) {
		if j >= 0 && j < len(numbers[i]) {
			if numbers[i][j] != 0 {
				num := numbers[i][j]

				scanJ := j
				for scanJ >= 0 && numbers[i][scanJ] == num {
					scanJ--
				}

				scanJ++

				for scanJ < len(numbers[i]) && numbers[i][scanJ] == num {
					numbers[i][scanJ] = 0
					scanJ++
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
	var currentNumber string
	var js []int
	for i, line := range engine {
		numbers[i] = make([]int, len(engine[i]))
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
	f, err := os.Create("numbers.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range numbers {
		for _, number := range line {
			f.WriteString(fmt.Sprintf("%d ", number))
		}
		f.WriteString("\n")
	}

	validNumbers := checkSymbols(engine, numbers)

	finalNumber := 0

	for _, num := range validNumbers {
		finalNumber += num
	}
	f2, err := os.Create("numbersafter.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range numbers {
		for _, number := range line {
			f2.WriteString(fmt.Sprintf("%d ", number))
		}
		f2.WriteString("\n")
	}

	fmt.Println(finalNumber)

}
