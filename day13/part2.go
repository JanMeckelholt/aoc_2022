package main

func part2(filePath string) (int, error) {
	pairs, err := readIn(filePath)
	if err != nil {
		return 0, err
	}

	ps := sortPackages(pairs)

	return findDividerPackages(ps), nil
}
