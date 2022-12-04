package main

func part2(path string) (score int, err error) {
	elvegroups, err := readIn(path)
	if err != nil {
		return 0, err
	}
	return countPartiallyContains(elvegroups), nil
}
