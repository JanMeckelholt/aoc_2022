package main

const size2 = 1000
const lengthSnake = 10

type position struct {
	row int
	col int
}

func part2(path string) (int, error) {
	input, err := readIn(path)
	if err != nil {
		return 0, err
	}
	grid := make([][]bool, 0)
	for i := 0; i < size2; i++ {
		row := make([]bool, size2)
		grid = append(grid, row)
	}
	snake := make([]*position, 0)
	for i := 0; i < lengthSnake; i++ {
		snake = append(snake, &position{row: size2 / 2, col: size2 / 2})
	}

	grid = *moveSnake(grid, snake, input)
	//printGrid(grid)
	return countVisited(grid), nil
}
