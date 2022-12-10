package main

func part2(path string) error {
	input, err := readIn(path)
	if err != nil {
		return err
	}

	cycles := make(map[int]int, 0)
	cycles = runCycles(cycles, input)

	write(cycles, 40)
	return nil
}
