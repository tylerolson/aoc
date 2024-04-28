package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var numberWords = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func replaceWords(line string) string {
	newLine := line

	for i := 0; i < len(numberWords); i++ {
		index := strings.Index(newLine, numberWords[i])
		for index != -1 {
			//since the strings can be overlapping, replace each word but leave the ends, three = t3e, so things like twone work two = t2o1e
			newLine = newLine[:index+1] + strconv.Itoa(i+1) + newLine[index+1:]
			index = strings.Index(newLine, numberWords[i])
		}
	}

	return newLine

}

func findNumber(line string) string {
	chars := "1234567890"

	first := string(line[strings.IndexAny(line, chars)])
	last := string(line[strings.LastIndexAny(line, chars)])

	return first + last
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	total := 0

	for scanner.Scan() {
		//part 2
		line := replaceWords(scanner.Text())

		number := findNumber(line)

		numint, err := strconv.Atoi(number)
		if err != nil {
			log.Fatal(err)
		}

		total += numint
		fmt.Println(scanner.Text())
		fmt.Printf("%d\n", numint)
	}

	fmt.Println("Final total:", total)
}
