package main

import (
	"fmt"
	"log"
)

func main() {

	topline, err := part1("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result: \nPart1: Top-Line: %s\n", topline)

	topline2, err := part2("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Part2 Top-Line: %s\n", topline2)

}
