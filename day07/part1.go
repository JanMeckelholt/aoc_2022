package main

func part1(path string) (int, error) {
	input, err := readIn(path)
	if err != nil {
		return 0, err
	}
	home := buildDirStructure(input)
	//showDir(&home)
	_, results := calcRecursiveDirSum(&home, 100000, 0, 0, []int{})
	sum := 0
	for _, subSum := range results {
		sum += subSum
	}
	return sum, nil
}
