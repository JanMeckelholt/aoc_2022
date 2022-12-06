package main

import (
	"fmt"
	"log"
)

func main() {

	num1, err := partSolution("./input.txt", 4)
	if err != nil {
		log.Fatal(err)
	}
	num2, err := partSolution("./input.txt", 14)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result: \nPart1: %d\n", num1)
	fmt.Printf("Part2: %d\n", num2)

}
