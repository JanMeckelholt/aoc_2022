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

	fmt.Printf("Result: \nPart1: %d visited positions\n", num1)

	num2example, err := part2("./example.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part2 example: %d visited positions\n", num2example)

	num2example2, err := part2("./example2.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part2 example2: %d visited positions\n", num2example2)

	num2, err := part2("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part2: %d visited positions\n", num2)

}
