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
	scanner := bufio.NewScanner(file)
	var inputdata []string
	for scanner.Scan() {
		inputdata = append(inputdata, scanner.Text())
	}
	binarycounts := getbinarycounts(inputdata)
	majority, minority := powerconsumption(binarycounts)
	gamma, _ := strconv.ParseInt(majority, 2, 64)
	epsilon, _ := strconv.ParseInt(minority, 2, 64)
	fmt.Println(gamma * epsilon)
	fmt.Println(lifesupportrating(inputdata))
}

func getbinarycounts(inputdata []string) []int {
	binarycounts := make([]int, 12)
	for n := 0; n < len(inputdata); n++ {
		binary_num := inputdata[n]
		for dig := 0; dig < 12; dig++ {
			if binary_num[dig] == '0' {
				binarycounts[dig] -= 1
			} else {
				binarycounts[dig] += 1
			}
		}
	}
	return binarycounts
}

func powerconsumption(binarycounts []int) (string, string) {
	var majority, minority string
	for n := 0; n < len(binarycounts); n++ {
		if binarycounts[n] < 0 {
			majority += "0"
			minority += "1"
		} else {
			majority += "1"
			minority += "0"
		}
	}
	return majority, minority
}

func lifesupportrating(inputdata []string) int64 {
	gen := make([]string, len(inputdata))
	copy(gen, inputdata)
	scrub := make([]string, len(inputdata))
	copy(scrub, inputdata)
	n := 0 // bit number 0 - 11
	x := 0 // row number 0 to length of data
	for n < 12 && len(gen) > 1 {
		binarycount := getbinarycounts(gen)
		maj, _ := powerconsumption(binarycount)
		x = 0
		for _, val := range gen {
			if maj[n] == val[n] {
				gen[x] = val
				x++
			}
		}
		gen = gen[:x]
		n++
	}
	fmt.Println("Gen", gen)
	n = 0
	for n < 12 && len(scrub) > 1 {
		binarycount_s := getbinarycounts(scrub)
		_, min := powerconsumption(binarycount_s)
		x = 0 // row value
		for _, val := range scrub {
			if min[n] == val[n] {
				scrub[x] = val
				x++
			}
		}
		scrub = scrub[:x]
		n++
	}
	fmt.Println("Scrub", scrub)
	ox, _ := strconv.ParseInt(gen[0], 2, 64)
	carb, _ := strconv.ParseInt(scrub[0], 2, 64)
	return ox * carb
}
