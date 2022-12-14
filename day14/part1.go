package main

func part1(filePath string) (int, error) {
	grid, err := readIn(filePath)
	if err != nil {
		return 0, err
	}
	//visualizeGrid(grid)
	return pourSand(grid, coordinate{x: inputX, y: 0}, heightPart1), nil
}
