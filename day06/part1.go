package main

func partSolution(path string, n int) (num int, err error) {
	input, err := readIn(path)
	if err != nil {
		return 0, err
	}
	return afterUnique(input, n), nil
}
