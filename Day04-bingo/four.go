package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	bingo, boards := loadinputdata()
	fmt.Println(bingo)
	fmt.Println(boards)
}

func loadinputdata() ([]string, [][][]int) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file")
		os.Exit(1)
	}
	defer file.Close()
	bingo := make([]string, 1)
	boards := make([][][]int, 200)
	scanner := bufio.NewScanner(file)
	line_num := 1
	row_index := 0
	board_index := 0
	boards[board_index] = make([][]int, 5)
	/* boards[board_index] => gives a full 5 by 5 board
	boards[board_index][row_index] => gives a row in a specific board
	boards[board_index][row_index][number] => gives a specific number
	*/

	for scanner.Scan() {
		singleline := scanner.Text()
		if line_num != 1 && singleline != "" {
			boards[board_index][row_index] = make([]int, 5)
			_, err := fmt.Sscanf(singleline, "%d %d %d %d %d", &boards[board_index][row_index][0], &boards[board_index][row_index][1], &boards[board_index][row_index][2], &boards[board_index][row_index][3], &boards[board_index][row_index][4])
			if err != nil {
				fmt.Println("Error", err)
			}
			row_index += 1
			if row_index == 5 {
				board_index += 1
				row_index = 0
				boards[board_index] = make([][]int, 5)
			}
		} else if line_num == 1 {
			bingo = append(bingo, scanner.Text())
			line_num += 1
		}
	}
	boards = boards[:board_index]
	return bingo, boards
}
