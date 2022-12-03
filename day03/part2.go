package main

func part2(path string) (score int, err error) {
	elveGroups, err := readInPart2(path)
	if err != nil {
		return 0, err
	}
	return getScorePart2(elveGroups), nil
}
