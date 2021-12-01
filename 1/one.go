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
	lastindex := len(inputarr) - 2
	incCount := 0
	for n := 0; n <= lastindex; n++ {
		if inputarr[n] < inputarr[n+1] {
			incCount++
		}
	}
	fmt.Println(incCount)
}
