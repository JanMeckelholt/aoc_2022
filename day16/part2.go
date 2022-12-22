package main

const minutesPart2 = 26

func part2(filePath string) (int, error) {
	caves, err := readIn(filePath)
	if err != nil {
		return 0, err
	}
	caves2 := rebuildMap(caves)
	caves2 = rebuildMapToIncludeOnlyValves(caves2)
	max := findMaxWithHelp(caves2, minutesPart2, start, start, 0, -1, -1, "", "")
	return max, nil
}
