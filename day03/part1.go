package main

func part1(path string) (score int, err error) {
	compartments, err := readIn(path)
	if err != nil {
		return 0, err
	}
	return getScore(compartments), nil
}
