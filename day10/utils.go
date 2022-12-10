package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const opAdd = "addx"
const opNil = "noop"

type operation struct {
	op  string
	num int
}

func readIn(path string) ([]operation, error) {
	var n int64
	input := make([]operation, 0)
	file, err := os.Open(path)
	if err != nil {

		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		inputString := scanner.Text()
		strs := strings.Split(inputString, " ")
		if len(strs) > 1 {
			n, err = strconv.ParseInt(strs[1], 10, 64)
			if err != nil {
				return input, err
			}
		} else {
			n = 0
		}
		op := operation{
			op:  strs[0],
			num: int(n),
		}
		input = append(input, op)
	}
	return input, nil
}

func operate(operation operation, cycletime, readingBefore int) (ct int, reading int) {
	if operation.op == opNil {
		return cycletime + 1, readingBefore
	}
	if operation.op == opAdd {
		return cycletime + 2, readingBefore + operation.num
	}
	fmt.Printf("\n%s", operation)
	log.Fatal("Wrong operation name!")
	return 0, 0
}

func runCycles(cycles map[int]int, operations []operation) map[int]int {
	ct := 1
	reading := 1
	cycles[0] = 1
	cycles[1] = 1
	for _, operation := range operations {
		if operation.op == opAdd {
			cycles[ct+1] = reading
		}
		ct, reading = operate(operation, ct, reading)
		cycles[ct] = reading
	}
	return cycles
}

func read(cycles map[int]int, readingPoints []int) int {
	var sum int
	for _, rP := range readingPoints {
		sum += cycles[rP] * rP
	}
	return sum
}

func write(cycles map[int]int, screenWidth int) {
	row := 0
	for i := 1; i < len(cycles); i++ {
		reading, ok := cycles[i]
		if !ok {
			log.Fatalf("no reading for cycle: %d", i)
		}
		if reading > i-(row*screenWidth)-3 && reading < i-(row*screenWidth)+1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
		if i%screenWidth == 0 {
			fmt.Print("\n")
			row++
		}
	}
}
