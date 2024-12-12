package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func day03part02() {
	mulRegex := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	doRegex := regexp.MustCompile(`do\(\)`)
	dontRegex := regexp.MustCompile(`don't\(\)`)

	file, err := os.Open("./input.txt")
	check(err)
	scanner := bufio.NewScanner(file)
	var allLines string
	var linesArray []string

	for scanner.Scan() {
		line := scanner.Text()
		allLines += line
	}
	linesArray = strings.Split(allLines, ")")

	var total int
	var enabled bool = true

	for i := 0; i < len(linesArray); i++ {
		linesArray[i] = linesArray[i] + ")"
		itsMul := mulRegex.FindAllString(linesArray[i], 1)
		itsDo := doRegex.FindAllString(linesArray[i], 1)
		itsDont := dontRegex.FindAllString(linesArray[i], 1)
		if itsDo != nil {
			enabled = true
		} else if itsDont != nil {
			enabled = false
		} else if itsMul != nil {
			if enabled {
				n1, n2 := getNumbers(itsMul[0])
				total += n1 * n2
			}
		}
	}
	fmt.Println("part 02:", total)
}
