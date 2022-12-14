package main

func part1(filePath string) (int, error) {
	grid, start, end, err := readIn(filePath)
	if err != nil {
		return 0, err
	}
	ways := make([]int, 0)
	path := make([]coordinate, 0)

	ways, _ = findShortestWay(grid, start, end, 0, ways, path)

	return getShortestWay(ways), nil
}
