package main

func part1(path string) (int, error) {
	input, err := readIn(path)
	if err != nil {
		return 0, err
	}

	cycles := make(map[int]int, 0)
	cycles = runCycles(cycles, input)

	return read(cycles, cycleReadingsPart1), nil
}
