package main

import (
	"bufio"
	"fmt"
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

func checkSymbols(engine [][]rune, numbers [][]int) ([]int, []int) {
	var validNumbers []int
	var gearRatios []int
	for i, line := range engine {
		for j, letter := range line {
			if !unicode.IsDigit(letter) && string(letter) != "." {
				var numbersAroundSymbol []int
				if value := isNumberAtIndex(numbers, i-1, j-1); value != -1 {
					numbersAroundSymbol = append(numbersAroundSymbol, value)
				}

				if value := isNumberAtIndex(numbers, i, j-1); value != -1 {
					numbersAroundSymbol = append(numbersAroundSymbol, value)
				}

				if value := isNumberAtIndex(numbers, i+1, j-1); value != -1 {
					numbersAroundSymbol = append(numbersAroundSymbol, value)
				}

				if value := isNumberAtIndex(numbers, i-1, j); value != -1 {
					numbersAroundSymbol = append(numbersAroundSymbol, value)
				}

				if value := isNumberAtIndex(numbers, i+1, j); value != -1 {
					numbersAroundSymbol = append(numbersAroundSymbol, value)
				}

				if value := isNumberAtIndex(numbers, i-1, j+1); value != -1 {
					numbersAroundSymbol = append(numbersAroundSymbol, value)
				}

				if value := isNumberAtIndex(numbers, i, j+1); value != -1 {
					numbersAroundSymbol = append(numbersAroundSymbol, value)
				}

				if value := isNumberAtIndex(numbers, i+1, j+1); value != -1 {
					numbersAroundSymbol = append(numbersAroundSymbol, value)
				}

				if string(letter) == "*" && len(numbersAroundSymbol) == 2 { //gear ratio
					gearRatio := 1
					for _, num := range numbersAroundSymbol {
						gearRatio *= num
					}

					gearRatios = append(gearRatios, gearRatio)
				}
				validNumbers = append(validNumbers, numbersAroundSymbol...)
			}
		}
	}

	return validNumbers, gearRatios
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
			}
			// IT DOESNT KNOW IF THE NUMBER IS DONE UNTIL NEXT ITERATION IF ITS ON THE END
			// SO WE NEED TO ALSO CHECK IF ITS ON THE END OR IT GETS PUT IN AT I+1
			if !unicode.IsDigit(letter) || j == len(engine[i])-1 {
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
	validNumbers, gearRatios := checkSymbols(engine, numbers)

	finalNumber := 0
	finalRatio := 0
	for _, num := range validNumbers {
		finalNumber += num
	}

	for _, num := range gearRatios {
		finalRatio += num
	}

	fmt.Println(finalNumber)
	fmt.Println(finalRatio)
}
