package main

func part2(path string) (int, error) {
	input, err := readIn(path)
	if err != nil {
		return 0, err
	}
	return countVisibleScenic(input), nil
}
