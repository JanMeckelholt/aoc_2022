package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type position struct {
	height       int
	shortestPath int
}

type coordinate struct {
	x int
	y int
}

func (c coordinate) equals(c2 coordinate) bool {
	return c.x == c2.x && c.y == c2.y
}

const (
	up = iota
	down
	right
	left
)

const goalHeight = 123 - 96
const startHeight = 97 - 96

func readIn(path string) (grid map[int]map[int]*position, start, end coordinate, err error) {
	input := make(map[int]map[int]*position, 0)
	file, err := os.Open(path)
	if err != nil {

		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	rowIndex := 0
	for scanner.Scan() {
		inputString := strings.TrimSpace(scanner.Text())
		row := make(map[int]*position, 0)
		for i, c := range inputString {
			h, err := strconv.ParseInt(fmt.Sprintf("%d", c), 10, 64)
			if err != nil {
				return nil, start, end, err
			}
			p := &position{
				height: int(h) - 96,
			}
			if c == 'E' {
				end.x = i
				end.y = rowIndex
				p.height = goalHeight
			}
			if c == 'S' {
				start.x = i
				start.y = rowIndex
				p.height = startHeight
			}
			row[i] = p
		}
		input[rowIndex] = row
		rowIndex++
	}
	return input, start, end, nil
}

func findShortestWay(grid map[int]map[int]*position, start, end coordinate, steps int, ways []int, path []coordinate) ([]int, []coordinate) {
	newWays := ways
	sumPath := make([]coordinate, 0)
	sumPath = append(sumPath, path...)

	path = append(path, start)
	if grid[start.y][start.x].shortestPath == 0 || grid[start.y][start.x].shortestPath > steps {
		grid[start.y][start.x].shortestPath = steps
		for i := 0; i < 4; i++ {
			switch i {
			case up:
				if start.y <= 0 {
					continue
				}
				newCoordinate := start
				newCoordinate.y = start.y - 1
				if pathIncludes(path, newCoordinate) {
					continue
				}
				if grid[newCoordinate.y][newCoordinate.x].height-grid[start.y][start.x].height > 1 {
					continue
				}
				if end.equals(newCoordinate) {
					newWays = append(newWays, steps+1)
					fmt.Printf("\n%5d - %d", len(newWays), steps+1)
					break
				}
				//path = append(path, newCoordinate)
				newPath := make([]coordinate, 0)
				resultsBefore := len(newWays)
				newWays, newPath = findShortestWay(grid, newCoordinate, end, steps+1, newWays, sumPath)
				if resultsBefore == len(newWays) {
					sumPath = mergePathes(sumPath, newPath)
				}
			case down:
				if start.y >= len(grid)-1 {
					continue
				}
				newCoordinate := start
				newCoordinate.y = start.y + 1
				if pathIncludes(path, newCoordinate) {
					continue
				}
				if grid[newCoordinate.y][newCoordinate.x].height-grid[start.y][start.x].height > 1 {
					continue
				}
				if end.equals(newCoordinate) {
					newWays = append(newWays, steps+1)
					fmt.Printf("\n%5d - %d", len(newWays), steps+1)
					break
				}
				//path = append(path, newCoordinate)
				newPath := make([]coordinate, 0)
				resultsBefore := len(newWays)
				newWays, newPath = findShortestWay(grid, newCoordinate, end, steps+1, newWays, sumPath)
				if resultsBefore == len(newWays) {
					sumPath = mergePathes(sumPath, newPath)
				}
			case right:
				if start.x >= len(grid[start.y])-1 {
					continue
				}
				newCoordinate := start
				newCoordinate.x = start.x + 1
				if pathIncludes(path, newCoordinate) {
					continue
				}
				if grid[newCoordinate.y][newCoordinate.x].height-grid[start.y][start.x].height > 1 {
					continue
				}
				if end.equals(newCoordinate) {
					newWays = append(newWays, steps+1)
					fmt.Printf("\n%5d - %d", len(newWays), steps+1)
					break
				}
				//path = append(path, newCoordinate)
				newPath := make([]coordinate, 0)
				resultsBefore := len(newWays)
				newWays, newPath = findShortestWay(grid, newCoordinate, end, steps+1, newWays, sumPath)
				if resultsBefore == len(newWays) {
					sumPath = mergePathes(sumPath, newPath)
				}
			case left:
				if start.x <= 0 {
					continue
				}
				newCoordinate := start
				newCoordinate.x = start.x - 1
				if pathIncludes(path, newCoordinate) {
					continue
				}
				if grid[newCoordinate.y][newCoordinate.x].height-grid[start.y][start.x].height > 1 {
					continue
				}
				if end.equals(newCoordinate) {
					newWays = append(newWays, steps+1)
					break
				}
				//path = append(path, newCoordinate)
				newPath := make([]coordinate, 0)
				resultsBefore := len(newWays)
				newWays, newPath = findShortestWay(grid, newCoordinate, end, steps+1, newWays, sumPath)
				if resultsBefore == len(newWays) {
					sumPath = mergePathes(sumPath, newPath)
				}
			default:
				log.Fatal("direction unknown!")
			}
		}
	}
	return newWays, sumPath
}

func getShortestWay(ways []int) int {
	if len(ways) == 0 {
		return 0
	}
	shortestWay := ways[0]
	for _, way := range ways {
		if way < shortestWay {
			shortestWay = way
		}
	}
	return shortestWay
}

func pathIncludes(path []coordinate, coordinate coordinate) bool {
	for _, c := range path {
		if c.x == coordinate.x && c.y == coordinate.y {
			return true
		}
	}
	return false
}

func mergePathes(path1, path2 []coordinate) (sumPath []coordinate) {
	sumPath = append(sumPath, path1...)
	for _, coordinate := range path2 {
		if !pathIncludes(sumPath, coordinate) {
			sumPath = append(sumPath, coordinate)
		}
	}
	return sumPath
}

func visualizePath(path []coordinate, grid map[int]map[int]*position) {
	var maxX, maxY int
	for _, c := range path {
		if c.x > maxX {
			maxX = c.x
		}
		if c.y > maxY {
			maxY = c.y
		}
	}
	for rowIndex := 0; rowIndex < maxY; rowIndex++ {
		fmt.Print("\n")
		for i := 0; i < maxX+1; i++ {
			if pathIncludes(path, coordinate{x: i, y: rowIndex}) {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
			fmt.Printf("%2d ", grid[rowIndex][i].height)
		}

	}
	fmt.Printf("\n")

}

func findAs(grid map[int]map[int]*position) (coordinates []coordinate) {
	coordinates = make([]coordinate, 0)
	for rowIndex, row := range grid {
		for i, pos := range row {
			if pos.height == 1 {
				coordinates = append(coordinates, coordinate{x: i, y: rowIndex})
			}
		}
	}
	return coordinates
}

func findShortestWayFromAnyA(grid map[int]map[int]*position, end coordinate) int {
	possibleStarts := findAs(grid)
	var minSteps int
	for _, possibleStart := range possibleStarts {
		ways := make([]int, 0)
		path := make([]coordinate, 0)
		ways, path = findShortestWay(grid, possibleStart, end, 0, ways, path)
		if len(ways) > 0 {
			s := getShortestWay(ways)
			if minSteps == 0 || minSteps > s {
				minSteps = s
			}
		}

	}
	return minSteps
}
