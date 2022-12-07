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

	num2, err := part2("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: \nPart1: %d\n", num1)
	fmt.Printf("Part2: %d\n", num2)

}
