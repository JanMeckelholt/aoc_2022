package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type movement struct {
	direction rune
	num       int
}

func readIn(path string) ([]movement, error) {
	input := make([]movement, 0)
	file, err := os.Open(path)
	if err != nil {

		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		inputString := scanner.Text()
		strs := strings.Split(inputString, " ")
		n, err := strconv.ParseInt(strs[1], 10, 64)
		if err != nil {
			return input, err
		}
		m := movement{
			direction: rune(strs[0][0]),
			num:       int(n),
		}
		input = append(input, m)
	}
	return input, nil
}

func move(grid [][]bool, rowHead, colHead, rowTail, colTail int, movements []movement) *[][]bool {
	for _, m := range movements {
		for i := 0; i < m.num; i++ {
			switch m.direction {
			case 'R':
				colHead++
			case 'L':
				colHead--
			case 'U':
				rowHead++
			case 'D':
				rowHead--
			}
			rowTail, colTail = followHead(rowHead, colHead, rowTail, colTail)
			grid[rowTail][colTail] = true
		}
	}
	return &grid
}

func followHead(rowHead, colHead, rowTail, colTail int) (newRowTail int, newColTail int) {
	if rowHead-rowTail > 1 {
		rowTail++
		colTail = colHead
	}
	if rowHead-rowTail < -1 {
		rowTail--
		colTail = colHead
	}
	if colHead-colTail > 1 {
		colTail++
		rowTail = rowHead
	}
	if colHead-colTail < -1 {
		colTail--
		rowTail = rowHead
	}
	return rowTail, colTail
}

func countVisited(grid [][]bool) int {
	var sum int
	for _, row := range grid {
		for _, v := range row {
			if v {
				sum++
			}
		}
	}
	return sum
}

func moveSnake(grid [][]bool, snake []*position, movements []movement) *[][]bool {
	for _, m := range movements {
		for j := 0; j < m.num; j++ {
			for i, _ := range snake {
				if i == 0 {
					switch m.direction {
					case 'R':
						snake[i].col++
					case 'L':
						snake[i].col--
					case 'U':
						snake[i].row++
					case 'D':
						snake[i].row--
					}
				} else {
					snake[i].row, snake[i].col = followHeadSnake(snake[i-1].row, snake[i-1].col, snake[i].row, snake[i].col)
					if i == len(snake)-1 {
						grid[snake[i].row][snake[i].col] = true
					}
				}
			}
		}
	}
	return &grid
}

func printGrid(grid [][]bool) {
	var minRow, maxRow, minCol, maxCol int
	minRow = len(grid[0]) - 1
	minCol = len(grid) - 1
	for rowIndex, row := range grid {
		for colIndex, v := range row {
			if v {
				if minRow > rowIndex {
					minRow = rowIndex
				}
				if maxRow < rowIndex {
					maxRow = rowIndex
				}
				if minCol > colIndex {
					minCol = colIndex
				}
				if maxCol < colIndex {
					maxCol = colIndex
				}
			}
		}
	}
	for i := minRow; i <= maxRow; i++ {
		for j := minCol; j <= maxCol; j++ {
			if grid[i][j] {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	fmt.Printf("\n%d - %d - %d - %d", minRow, maxRow, minCol, maxCol)
}

func followHeadSnake(rowHead, colHead, rowTail, colTail int) (newRowTail int, newColTail int) {
	if rowHead-rowTail > 1 {
		rowTail++
		switch colHead - colTail {
		case 2:
			colTail++
		case -2:
			colTail--
		default:
			colTail = colHead
		}

	}
	if rowHead-rowTail < -1 {
		rowTail--
		switch colHead - colTail {
		case 2:
			colTail++
		case -2:
			colTail--
		default:
			colTail = colHead
		}
	}
	if colHead-colTail > 1 {
		colTail++
		switch rowHead - rowTail {
		case 2:
			rowTail++
		case -2:
			rowTail--
		default:
			rowTail = rowHead
		}

	}
	if colHead-colTail < -1 {
		colTail--
		switch rowHead - rowTail {
		case 2:
			rowTail++
		case -2:
			rowTail--
		default:
			rowTail = rowHead
		}
	}
	return rowTail, colTail
}
