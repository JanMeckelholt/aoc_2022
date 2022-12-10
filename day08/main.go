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

	fmt.Printf("Result: \nPart1: %d visible trees", num1)

	num2, err := part2("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nPart2: %d max scenic count\n", num2)

}
