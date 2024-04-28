package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findNumber(line string) string {
	strSlice := strings.Split(line, "")
	f := func(r rune) bool {
		return r == '1' || r == '2' || r == '3' || r == '4' || r == '5' || r == '6' || r == '7' || r == '8' || r == '9'
	}

	first := ""
	last := ""

	for i := 0; i < len(strSlice); i++ {
		if strings.ContainsFunc(strSlice[i], f) {
			if first == "" {
				first = strSlice[i]
			}
			last = strSlice[i]
		}
	}

	return first + last
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	total := 0

	for scanner.Scan() {
		number := findNumber(scanner.Text())

		numint, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}

		total += numint
		fmt.Println(scanner.Text())
		fmt.Printf("%d\n", numint)
	}

	fmt.Println("Final total:", total)
}
