package main

func part2(filePath string) (int, error) {
	grid, err := readIn(filePath)
	if err != nil {
		return 0, err
	}
	h := getFloor(grid)
	return pourSandPart2(grid, coordinate{x: inputX, y: 0}, h), nil
}
