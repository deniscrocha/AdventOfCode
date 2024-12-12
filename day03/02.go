package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func day03part02() {
	mulRegex := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	doRegex := regexp.MustCompile(`do\(\)`)
	dontRegex := regexp.MustCompile(`don't\(\)`)

	file, err := os.Open("./test2.txt")
	check(err)
	scanner := bufio.NewScanner(file)
	total := 0
	var matches []string
	var allLines string

	for scanner.Scan() {
		line := scanner.Text()
		allLines += line
	}

	matches = mulRegex.FindAllString(allLines, -1)

	matchesIndex := mulRegex.FindAllStringIndex(allLines, -1)

	var indexes []int

	dont := dontRegex.FindAllStringIndex(allLines, -1)
	do := doRegex.FindAllStringIndex(allLines, -1)


	for i := 0; i < len(matches); i++ {
		text := matches[i]
		lhs, rhs := getNumbers(text)
		total += lhs * rhs
	}

	fmt.Println("part 02:", total)
}