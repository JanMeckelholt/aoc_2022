package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	x int
	y int
}

const width = 600
const heightPart1 = 600
const inputX = 500

func readIn(path string) (map[int]map[int]bool, error) {
	var start, end coordinate
	input := make(map[int]map[int]bool)
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for i := 0; i < heightPart1; i++ {
		row := make(map[int]bool)
		input[i] = row
	}

	for scanner.Scan() {
		inputString := strings.TrimSpace(scanner.Text())
		strs := strings.Split(inputString, " -> ")
		for i, s := range strs {
			if i < len(strs)-1 {
				start = strToCoordinate(s)
				end = strToCoordinate(strs[i+1])
				input = addRock(input, start, end)
			}
		}
	}
	return input, nil
}

func strToCoordinate(s string) coordinate {
	strs := strings.Split(s, ",")
	x, _ := strconv.ParseInt(strs[0], 10, 64)
	y, _ := strconv.ParseInt(strs[1], 10, 64)
	return coordinate{x: int(x), y: int(y)}
}

func addRock(grid map[int]map[int]bool, start, end coordinate) map[int]map[int]bool {
	var s, b int
	if start.x == end.x {
		if start.y > end.y {
			s = end.y
			b = start.y
		} else {
			s = start.y
			b = end.y
		}
		for i := s; i <= b; i++ {
			grid[i][start.x] = true
		}
	}
	if start.y == end.y {
		if start.x > end.x {
			s = end.x
			b = start.x
		} else {
			s = start.x
			b = end.x
		}
		for i := s; i <= b; i++ {
			grid[start.y][i] = true
		}
	}
	return grid
}

func getFloor(grid map[int]map[int]bool) int {
	var floor int
	for rowIndex, row := range grid {
		for _, c := range row {
			if c {
				if rowIndex > floor {
					floor = rowIndex
					break
				}
			}
		}
	}

	return floor + 2
}

func visualizeGrid(grid map[int]map[int]bool, height int) {
	for rowIndex := 0; rowIndex < height; rowIndex++ {
		fmt.Print("\n")
		for i := 0; i < width; i++ {
			if grid[rowIndex][i] {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}

	}
	fmt.Printf("\n")

}

func pourSand(grid map[int]map[int]bool, outletX coordinate, height int) int {
	var numSand int
	var fallenIntoVoid bool

	for !fallenIntoVoid {
		numSand++
		pos := outletX
		for i := 0; i < height; i++ {
			if i == height-1 {
				fallenIntoVoid = true
				return numSand - 1
			}
			if grid[i+1][pos.x] {
				if !grid[i+1][pos.x-1] {
					pos.x = pos.x - 1
				} else if !grid[i+1][pos.x+1] {
					pos.x = pos.x + 1
				} else {
					grid[i][pos.x] = true
					break
				}
			}
		}
	}
	return numSand
}

func pourSandPart2(grid map[int]map[int]bool, outletX coordinate, height int) int {
	var numSand int
	var reachedOutlet bool

	for !reachedOutlet {
		numSand++
		pos := outletX
		for i := 0; i < height; i++ {
			if i == height-1 {
				grid[i][pos.x] = true
			}
			if grid[i+1][pos.x] {
				if !grid[i+1][pos.x-1] {
					pos.x = pos.x - 1
				} else if !grid[i+1][pos.x+1] {
					pos.x = pos.x + 1
				} else {
					grid[i][pos.x] = true
					if i == 0 {
						return numSand
					}
					break
				}
			}
		}
	}
	return numSand
}
