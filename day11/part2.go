package main

import "fmt"

func part2(path string) (int, error) {
	monkies, err := readIn(path)
	if err != nil {
		return 0, err
	}

	for i := 0; i < 10000; i++ {
		monkies = inspectionRoundPart2(monkies)
	}

	for _, m := range monkies {
		fmt.Printf("\n Monkey %d - %d", m.num, m.numInspections)
	}

	return getMonkeyBusiness(monkies), nil
}
