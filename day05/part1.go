package main

func part1(path string) (topline string, err error) {
	stacks, movements, err := readIn(path)
	if err != nil {
		return "", err
	}
	//showStacks(stacks)
	//showMovements(movements)
	stacks = moveAll(stacks, movements)
	//showStacks(stacks)
	return readTopLine(stacks), nil
}
