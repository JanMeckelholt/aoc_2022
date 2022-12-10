package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type lineType int

func readIn(path string) ([][]int, error) {
	var value int64
	input := make([][]int, 0)
	file, err := os.Open(path)
	if err != nil {

		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		inputString := scanner.Text()
		row := make([]int, 0)
		for _, v := range inputString {
			value, err = strconv.ParseInt(string(v), 10, 64)
			if err != nil {
				return input, err
			}
			row = append(row, int(value))
		}
		input = append(input, row)
	}
	return input, nil
}

func checkRowFromLeft(row []int, index int) bool {
	for i, v := range row {
		if i == index {
			return true
		}
		if v >= row[index] {
			return false
		}
	}
	return false
}

func checkRowFromRight(row []int, index int) bool {
	for i := len(row) - 1; i >= index; i-- {
		if i == index {
			return true
		}
		if row[i] >= row[index] {
			return false
		}
	}
	return false
}

func checkCol(grid [][]int, col, index int) bool {
	column := make([]int, 0)
	for i := 0; i < len(grid); i++ {
		column = append(column, grid[i][col])
	}
	return checkRowFromRight(column, index) || checkRowFromLeft(column, index)
}

func countVisible(grid [][]int) int {
	var sum int
	for rowIndex, row := range grid {
		for colIndex, _ := range grid[rowIndex] {
			if checkRowFromRight(row, colIndex) || checkRowFromLeft(row, colIndex) || checkCol(grid, colIndex, rowIndex) {
				sum++
			}
		}
	}
	return sum
}

func countToRight(row []int, index int) int {
	count := 0
	if index == len(row)-1 {
		return 0
	}
	for i := index + 1; i < len(row); i++ {
		count++
		if i == len(row)-1 {
			return count
		}
		if row[i] >= row[index] {
			return count
		}
	}
	return count
}

func countToLeft(row []int, index int) int {
	count := 0
	if index == 0 {
		return 0
	}
	for i := index - 1; i >= 0; i-- {
		count++
		if i == 0 {
			return count
		}
		if row[i] >= row[index] {
			return count
		}
	}
	return count
}

func countCol(grid [][]int, col, index int) int {
	column := make([]int, 0)
	for i := 0; i < len(grid); i++ {
		column = append(column, grid[i][col])
	}
	return countToLeft(column, index) * countToRight(column, index)
}

func countVisibleScenic(grid [][]int) int {
	var max int
	for rowIndex, row := range grid {
		for colIndex, _ := range row {
			if countToRight(row, colIndex)*countToLeft(row, colIndex)*countCol(grid, colIndex, rowIndex) > max {
				max = countToRight(row, colIndex) * countToLeft(row, colIndex) * countCol(grid, colIndex, rowIndex)
			}
		}
	}
	return max
}
