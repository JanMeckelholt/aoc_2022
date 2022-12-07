package main

import "fmt"

func part2(path string) (int, error) {
	input, err := readIn(path)
	if err != nil {
		return 0, err
	}
	home := buildDirStructure(input)
	//showDir(&home)
	_, results := calcRecursiveDirSum(&home, 10000000000000000, 0, 0, []int{})
	total := results[len(results)-1]
	spaceToBeFreedUp := total + 30000000 - 70000000
	fmt.Printf("\nspace to be freed up: %d\n", spaceToBeFreedUp)
	minSufficient := total
	for _, subSum := range results {
		if subSum < minSufficient && subSum > spaceToBeFreedUp {
			minSufficient = subSum
		}
	}
	return minSufficient, nil
}
