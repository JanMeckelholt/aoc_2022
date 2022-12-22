package main

import (
	"fmt"
	"log"
)

func main() {

	num1Example, err := part1("./example.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result: \nPart1 Example: max %d can be relased in %d minutes.\n", num1Example, minutes)

	num1, err := part1("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Part1: max %d can be relased in %d minutes.\n", num1, minutes)

	num2Example, err := part2("./example.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part2 Example: max %d can be relased in %d minutes.\n", num2Example, minutesPart2)

	num2, err := part2("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part2: max %d can be relased in %d minutes.\n", num2, minutesPart2)
}
