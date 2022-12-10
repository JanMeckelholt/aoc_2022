package main

const size = 1000

func part1(path string) (int, error) {
	input, err := readIn(path)
	if err != nil {
		return 0, err
	}
	grid := make([][]bool, 0)
	for i := 0; i < size; i++ {
		row := make([]bool, size)
		grid = append(grid, row)
	}
	grid = *move(grid, size/2, size/2, size/2, size/2, input)

	return countVisited(grid), nil
}
