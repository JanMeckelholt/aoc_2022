package main

func part1(path string) (int, error) {
	input, err := readIn(path)
	if err != nil {
		return 0, err
	}
	return countVisible(input), nil
}
