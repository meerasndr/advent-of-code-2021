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
	var binarycounts [12]int
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
	var g, e string
	for n := 0; n < len(binarycounts); n++ {
		if binarycounts[n] < 0 {
			g += "0"
			e += "1"
		} else {
			g += "1"
			e += "0"
		}
	}
	gamma, _ := strconv.ParseInt(g, 2, 64)
	epsilon, _ := strconv.ParseInt(e, 2, 64)
	fmt.Println(gamma * epsilon)
}
