package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var leftList = &LinkedList{head: nil, length: 0}
var rightList = &LinkedList{head: nil, length: 0}

func main() {
	file, err := os.Open("./input01.txt")
	check(err)

	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		line := scanner.Text()
		result := strings.Split(line, "   ")
		firstNum, err := strconv.Atoi(result[0])
		check(err)
		leftList.insertAtBack(firstNum)
		secondNum, err := strconv.Atoi(result[1])
		check(err)
		rightList.insertAtBack(secondNum)
	}
	leftList.bubbleSort()
	rightList.bubbleSort()
	leftCurrent := leftList.head
	rightCurrent := rightList.head
	totalDif := 0
	for leftCurrent != nil{
		dif := leftCurrent.data - rightCurrent.data
		dif = int(math.Abs(float64(dif)))
		totalDif += dif
		leftCurrent = leftCurrent.next
		rightCurrent = rightCurrent.next
	}
	fmt.Println(totalDif)
}
