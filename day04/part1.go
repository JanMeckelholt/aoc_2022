package main

func part1(path string) (score int, err error) {
	elvegroups, err := readIn(path)
	if err != nil {
		return 0, err
	}
	return countFullyContains(elvegroups), nil
}
