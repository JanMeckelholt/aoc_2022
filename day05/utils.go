package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	crates []rune
}

type Movement struct {
	numberOfCrates int
	from           int
	to             int
}

func readIn(path string) (map[int]Stack, []Movement, error) {
	file, err := os.Open(path)
	if err != nil {

		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var (
		numberOfStacks int
	)
	stacks := make(map[int]Stack, 0)
	movements := make([]Movement, 0)
	firstPartOfInput := true

	for scanner.Scan() {
		inputString := scanner.Text()

		if firstPartOfInput { // Stacks
			if len(strings.TrimSpace(inputString)) != 0 {
				if numberOfStacks == 0 {
					numberOfStacks = (len(inputString) + 1) / 4
				}

				for i := 0; i < numberOfStacks; i++ {
					if inputString[i*4] == '[' {
						stack := stacks[i]
						stack.crates = append([]rune{rune(inputString[i*4+1])}, stack.crates...)
						stacks[i] = stack
					}
				}
			} else {
				firstPartOfInput = false
			}
		} else { //Movements
			inputString = inputString[5:]
			strs1 := strings.Split(inputString, " from ")
			strs2 := strings.Split(strs1[1], " to ")
			nOC, err := strconv.ParseInt(strs1[0], 10, 64)
			if err != nil {
				return stacks, movements, err
			}
			from, err := strconv.ParseInt(strs2[0], 10, 64)
			if err != nil {
				return stacks, movements, err
			}
			to, err := strconv.ParseInt(strs2[1], 10, 64)
			if err != nil {
				return stacks, movements, err
			}
			movement := Movement{
				numberOfCrates: int(nOC),
				from:           int(from) - 1,
				to:             int(to) - 1,
			}
			movements = append(movements, movement)
		}

	}
	return stacks, movements, nil
}

func singleMove(stacks map[int]Stack, movement Movement) map[int]Stack {
	f := stacks[movement.from]
	t := stacks[movement.to]
	for step := 0; step < movement.numberOfCrates; step++ {
		crate := f.crates[len(f.crates)-1]
		t.crates = append(t.crates, crate)
		f.crates = f.crates[:len(f.crates)-1]
	}
	stacks[movement.from] = f
	stacks[movement.to] = t
	return stacks
}

func compoundMove(stacks map[int]Stack, movement Movement) map[int]Stack {
	f := stacks[movement.from]
	t := stacks[movement.to]
	movePackage := f.crates[len(f.crates)-movement.numberOfCrates:]
	t.crates = append(t.crates, movePackage...)
	f.crates = f.crates[:len(f.crates)-movement.numberOfCrates]
	stacks[movement.from] = f
	stacks[movement.to] = t
	return stacks
}

func moveAll(stacks map[int]Stack, movements []Movement) map[int]Stack {
	for _, m := range movements {
		stacks = singleMove(stacks, m)
	}
	return stacks
}

func moveAllCompound(stacks map[int]Stack, movements []Movement) map[int]Stack {
	for _, m := range movements {
		stacks = compoundMove(stacks, m)
	}
	return stacks
}

func readTopLine(stacks map[int]Stack) string {
	var topLine string
	for i := 0; i < len(stacks); i++ {
		topLine += fmt.Sprintf("%c", stacks[i].crates[len(stacks[i].crates)-1])
	}
	return topLine
}

func showStacks(stacks map[int]Stack) {
	for i, stack := range stacks {
		fmt.Printf("\n%d", i)
		for _, r := range stack.crates {
			fmt.Printf(" %c ", r)
		}
		fmt.Println("\n----")
	}
}

func showMovements(movements []Movement) {
	for i, movement := range movements {
		fmt.Printf("\n%d", i)
		fmt.Printf(" num %d", movement.numberOfCrates)
		fmt.Printf(" from %d", movement.from)
		fmt.Printf(" to %d", movement.to)
		fmt.Println("\n----")
	}
}
