package main

func part2(path string) (score int, err error) {
	moves, err := readIn(path)
	if err != nil {
		return 0, err
	}
	moves = updateMoves(moves)
	return calcMatchScore(moves), nil
}
