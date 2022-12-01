package main

import (
	"fmt"
	"log"
)

func main() {

	maxElve, err := part1("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result: \nPart1: Max: %v\n", maxElve)

	maxTopThree, err := part2("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result: \nPart2: Max Top 3: %v\n", maxTopThree)

}
