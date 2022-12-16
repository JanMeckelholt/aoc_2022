package main

import (
	"fmt"
	"log"
)

func main() {
	num1Example, err := part1Example("./example.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result: \nPart1 Example: %d positions can not contain a beacon\n", num1Example)

	num1, err := part1("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: \nPart1: %d positions can not contain a beacon\n", num1)

	num2, err := part2("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part2: tuning frequency: %d \n", num2)

}
