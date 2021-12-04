package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file")
		os.Exit(1)
	}
	defer file.Close()
	var inputarr []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		x, _ := strconv.Atoi(scanner.Text())
		inputarr = append(inputarr, x)
	}
	inputlastindex := len(inputarr) - 2
	fmt.Println("Question 1a:", getIncreaseCount(inputarr, inputlastindex))

	var slidingwindow []int
	slidinglastindex := len(inputarr) - 3
	for n := 0; n <= slidinglastindex; n++ {
		slidingwindow = append(slidingwindow, inputarr[n]+inputarr[n+1]+inputarr[n+2])
	}
	fmt.Println("Question 1b:", getIncreaseCount(slidingwindow, len(slidingwindow)-2))
}

func getIncreaseCount(input []int, lastindex int) int {
	incCount := 0
	for n := 0; n <= lastindex; n++ {
		if input[n] < input[n+1] {
			incCount++
		}
	}
	return incCount
}
