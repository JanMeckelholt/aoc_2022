package main

func part2(path string) (maxTopThre int, err error) {
	valueMap, err := readIn(path)
	if err != nil {
		return 0, err
	}
	return getTopThreeSum(valueMap), nil
}
