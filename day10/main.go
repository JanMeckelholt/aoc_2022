package main

import (
	"fmt"
	"log"
)

var cycleReadingsPart1 = []int{20, 60, 100, 140, 180, 220}

func main() {
	num1, err := part1("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: \nPart1: %d visited positions\n", num1)

	fmt.Printf("\nPart2 Example:\n")
	err = part2("./example.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nPart2:\n")
	err = part2("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
}
