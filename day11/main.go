package main

import (
	"fmt"
	"log"
)

func main() {
	num1, err := part1("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: \nPart1: %d Monkeybusiness ", num1)

	num2, err := part2("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nPart2: %d Monkeybusiness ", num2)
}
