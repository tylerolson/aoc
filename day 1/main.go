package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	total := 0
	number := ""

	for scanner.Scan() {
		regex := regexp.MustCompile("\\d")
		arr := regex.FindAllString(scanner.Text(), -1)
		number = ""

		for index, element := range arr {
			if index == 0 || index == len(arr)-1 {
				if len(arr) == 1 {
					number += element
				}
				number += element
			}
		}

		numint, _ := strconv.Atoi(number)
		total += numint
		fmt.Println(scanner.Text())
		fmt.Println(number + "\n")
	}

	fmt.Println("Final total:", total)
}
