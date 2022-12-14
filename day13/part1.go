package main

func part1(filePath string) (int, error) {
	pairs, err := readIn(filePath)
	if err != nil {
		return 0, err
	}

	return checkInput(pairs), nil
}
