package main

func part2(filePath string) (int, error) {
	grid, _, end, err := readIn(filePath)
	if err != nil {
		return 0, err
	}

	return findShortestWayFromAnyA(grid, end), nil
}
