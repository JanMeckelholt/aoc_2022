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

	fmt.Printf("Result: \nPart1: Ordered sum is %d\n", num1)

	num2, err := part2("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nPart2: decoder key is %d\n", num2)
}
