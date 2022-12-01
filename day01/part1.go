package main

func part1(path string) (max int, err error) {
	valueMap, err := readIn(path)
	if err != nil {
		return 0, err
	}
	return getMax(valueMap), nil
}
