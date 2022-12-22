package main

const minutes = 30
const start = "AA"
const openingStr = "%%"

func part1(filePath string) (int, error) {
	caves, err := readIn(filePath)
	if err != nil {
		return 0, err
	}
	visited := make([]string, 0)
	caves2 := rebuildMap(caves)
	max := findMax2(caves2, minutes, start, 0, 0, visited, "")
	return max, nil
}
