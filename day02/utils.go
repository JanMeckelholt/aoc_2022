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
		mM := rune(moves[0][0])
		oM := rune(moves[1][0])
		valueMap = append(valueMap, Rps{myMove: mM, opponentsMove: oM})
	}
	return valueMap, nil
}
