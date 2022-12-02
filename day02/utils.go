package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Rps = struct {
	myMove        rune
	opponentsMove rune
}

func readIn(path string) ([]Rps, error) {
	file, err := os.Open(path)
	if err != nil {

		log.Fatal(err)
	}
	defer file.Close()
	valueMap := make([]Rps, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputString := strings.TrimSpace(scanner.Text())
		moves := strings.Split(inputString, " ")
		mM := rune(moves[1][0])
		oM := rune(moves[0][0])
		valueMap = append(valueMap, Rps{myMove: mM, opponentsMove: oM})
	}
	return valueMap, nil
}

func calcSingleScore(rps Rps) int {
	switch rps.opponentsMove {
	case 'A':
		if rps.myMove == 'X' { //tied
			return 1 + 3
		}
		if rps.myMove == 'Y' { //won
			return 2 + 6
		}
		return 3 + 0

	case 'B':
		if rps.myMove == 'Y' { //tied
			return 2 + 3
		}
		if rps.myMove == 'Z' { //won
			return 3 + 6
		}
		return 1 + 0

	case 'C':
		if rps.myMove == 'Z' { //tied
			return 3 + 3
		}
		if rps.myMove == 'X' { //won
			return 1 + 6
		}
		return 2 + 0
	}
	return 0
}

func calcMatchScore(moves []Rps) int {
	var sum int
	for _, rps := range moves {
		sum += calcSingleScore(rps)
	}
	return sum
}

func getMyMove(rps Rps) Rps {
	var mM rune
	switch rps.opponentsMove {
	case 'A':
		if rps.myMove == 'X' { //lose
			mM = 'Z'
			break
		}
		if rps.myMove == 'Y' { //draw
			mM = 'X'
			break
		}
		mM = 'Y'

	case 'B':
		if rps.myMove == 'X' { //lose
			mM = 'X'
			break
		}
		if rps.myMove == 'Y' { //draw
			mM = 'Y'
			break
		}
		mM = 'Z'

	case 'C':
		if rps.myMove == 'X' { //lose
			mM = 'Y'
			break
		}
		if rps.myMove == 'Y' { //draw
			mM = 'Z'
			break
		}
		mM = 'X'
	}
	return Rps{
		myMove:        mM,
		opponentsMove: rps.opponentsMove,
	}
}

func updateMoves(moves []Rps) []Rps {
	updateMoves := make([]Rps, 0)
	for _, rps := range moves {
		move := getMyMove(rps)
		updateMoves = append(updateMoves, move)
	}
	return updateMoves
}
