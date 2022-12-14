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

	fmt.Printf("Result: \nPart1: %d units of sand can come to rest", num1)

	num2, err := part2("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nPart2: %d units of sand can come to rest\n", num2)

}
