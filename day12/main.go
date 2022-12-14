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

	fmt.Printf("Result: \nPart1: Shortest way has %d steps ", num1)

	num2, err := part2("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nPart2: Shortest way from any a has %d steps\n", num2)
}
