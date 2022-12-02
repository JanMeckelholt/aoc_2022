package main

func part1(path string) (score int, err error) {
	moves, err := readIn(path)
	if err != nil {
		return 0, err
	}
	return calcMatchScore(moves), nil
}
