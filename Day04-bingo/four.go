package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	bingo := loadinputdata()
	fmt.Println(bingo)
}

func loadinputdata() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file")
		os.Exit(1)
	}
	defer file.Close()
	bingo := make([]string, 1)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bingo = append(bingo, scanner.Text())
		break
	}
	return bingo
}
