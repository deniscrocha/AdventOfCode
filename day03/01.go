package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func day03part01() {
	regex := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	file, err := os.Open("./input.txt")
	check(err)
	scanner := bufio.NewScanner(file)
	total := 0
	var matches []string
	for scanner.Scan() {
		lineMatch := regex.FindAllString(scanner.Text(), -1)
		matches = append(matches, lineMatch[:]...)
	}
	for i := 0; i < len(matches); i++ {
		text := matches[i]
		lhs, rhs := getNumbers(text)
		total += lhs * rhs
	}
	fmt.Println("part 01:", total)
}

func getNumbers(text string) (int, int) {
	regex := regexp.MustCompile(`([0-9]{1,3})`)
	numbers := regex.FindAllString(text, -1)
	n1, err := strconv.Atoi(numbers[0])
	check(err)
	n2, err := strconv.Atoi(numbers[1])
	check(err)
	return n1, n2
}
