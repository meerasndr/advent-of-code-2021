package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	horizontal int
	depth      int
	aim        int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file")
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var inputarr [][]string
	for scanner.Scan() {
		inputarr = append(inputarr, strings.Split(scanner.Text(), " "))
	}
	fmt.Println("Planned Course One:", plannedCourseOne(inputarr))
	fmt.Println("Planned Course Two:", plannedCourseTwo(inputarr))
}

func plannedCourseOne(inputarr [][]string) int {
	mycoords := coordinate{horizontal: 0, depth: 0}
	for n := 0; n < len(inputarr); n++ {
		val, _ := strconv.Atoi(inputarr[n][1])
		if inputarr[n][0] == "forward" {
			mycoords.horizontal += val
		} else if inputarr[n][0] == "down" {
			mycoords.depth += val
		} else {
			mycoords.depth -= val
		}
	}
	return mycoords.horizontal * mycoords.depth
}

func plannedCourseTwo(inputarr [][]string) int {
	mycoords := coordinate{horizontal: 0, depth: 0, aim: 0}
	for n := 0; n < len(inputarr); n++ {
		val, _ := strconv.Atoi(inputarr[n][1])
		if inputarr[n][0] == "forward" {
			mycoords.horizontal += val
			mycoords.depth += mycoords.aim * val
		} else if inputarr[n][0] == "down" {
			mycoords.aim += val
		} else {
			mycoords.aim -= val
		}
	}
	return mycoords.horizontal * mycoords.depth
}
