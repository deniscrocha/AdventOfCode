package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day01part02() {
	var leftList = &LinkedList{head: nil, length: 0}
	var rightList = &LinkedList{head: nil, length: 0}

	file, err := os.Open("./input01.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		result := strings.Split(line, "   ")
		firstNum, err := strconv.Atoi(result[0])
		check(err)
		leftList.insertAtBack(firstNum)
		secondNum, err := strconv.Atoi(result[1])
		check(err)
		rightList.insertAtBack(secondNum)
	}
	file.Close()

	total := 0
	current := leftList.head
	for current != nil {
		number := current.data
		times := rightList.timesIn(number)
		if times == 0 {
			current = current.next
			continue
		}
		fmt.Println(number, " ",times)
		total += number * times
		current = current.next
	}
	fmt.Println(total)
}
