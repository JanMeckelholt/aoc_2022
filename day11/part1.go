package main

func part1(path string) (int, error) {
	monkies, err := readIn(path)
	if err != nil {
		return 0, err
	}

	for i := 0; i < 10000; i++ {
		monkies = inspectionRound(monkies)
	}

	return getMonkeyBusiness(monkies), nil
}
