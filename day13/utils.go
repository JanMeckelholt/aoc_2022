package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type inputStrs struct {
	left  string
	right string
}

const (
	empty = iota
	number
	group
)

const (
	ordered = iota
	unordered
	equal
)

var dividerPackages = []string{"[[2]]", "[[6]]"}
var debug int

func readIn(path string) ([]inputStrs, error) {
	input := make([]inputStrs, 0)
	file, err := os.Open(path)
	if err != nil {

		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var (
		i    int
		pair inputStrs
	)
	for scanner.Scan() {
		inputString := strings.TrimSpace(scanner.Text())
		i++
		switch i % 3 {
		case 1:
			pair.left = inputString
		case 2:
			pair.right = inputString
			input = append(input, pair)
		case 0:
			pair = inputStrs{}
		}
	}
	return input, nil
}

func strsAreOrdered(pair inputStrs) (int, inputStrs) {
	//fmt.Printf("\nBEGIN CHECK: %s against %s\n", pair.left, pair.right)
	elementLeft, remainderLeft := getFirstElement(pair.left)
	elementRight, remainderRight := getFirstElement(pair.right)
analyseElement:
	for len(elementLeft) > 0 {
		//	fmt.Printf("\nCheck: eL: %s - rL: %s || eR: %s - rR: %s", elementLeft, remainderLeft, elementRight, remainderRight)
		if len(elementRight) == 0 {
			return unordered, inputStrs{}
		}
		switch analyseElement(elementLeft) {
		case empty:
			switch analyseElement(elementRight) {
			case empty:
				return equal, inputStrs{left: remainderLeft, right: remainderRight}
			case number, group:
				return ordered, inputStrs{}
			}
		case number:
			switch analyseElement(elementRight) {
			case empty:
				return unordered, inputStrs{}
			case group:
				if elementRight[0] != '[' {
					nR, _ := strconv.ParseInt(elementRight[:strings.Index(elementRight, ",")], 10, 64)
					nL, _ := strconv.ParseInt(elementLeft, 10, 64)
					if nR >= nL {
						return ordered, inputStrs{}
					}
					return unordered, inputStrs{}
				}
				elementLeft = fmt.Sprintf("[%s]", elementLeft)
				continue analyseElement
			case number:
				nL, _ := strconv.ParseInt(elementLeft, 10, 64)
				nR, _ := strconv.ParseInt(elementRight, 10, 64)
				if nL > nR {
					return unordered, inputStrs{}
				}
				if nL < nR {
					return ordered, inputStrs{}
				}
				return equal, inputStrs{left: remainderLeft, right: remainderRight}
			}
		case group:
			switch analyseElement(elementRight) {
			case empty:
				return unordered, inputStrs{}
			case number:
				if elementLeft[0] != '[' {
					nL, _ := strconv.ParseInt(elementLeft[:strings.Index(elementLeft, ",")], 10, 64)
					nR, _ := strconv.ParseInt(elementRight, 10, 64)
					if nR > nL {
						return ordered, inputStrs{}
					}
					return unordered, inputStrs{}
				}
				elementRight = fmt.Sprintf("[%s]", elementRight)
				continue analyseElement
			case group:
				res, remainder := strsAreOrdered(inputStrs{left: elementLeft, right: elementRight})
				if res != equal {
					return res, inputStrs{}
				}
				if len(remainder.left) == 0 && len(remainder.right) != 0 {
					return ordered, inputStrs{}
				}
				elementLeft = remainder.left
				elementRight = remainder.right
				remainder = inputStrs{}

			}

		}
	}
	return equal, inputStrs{left: remainderLeft, right: remainderRight}
}

func getFirstElement(input string) (element string, remainder string) {
	if len(input) == 0 {
		return "", ""
	}
	if input[0] == '[' {
		if findClosingIndex(input) == len(input)-1 {
			return input[1 : len(input)-1], ""
		}
		return input[:findClosingIndex(input)+1], input[findClosingIndex(input)+2:]
	}
	if !strings.Contains(input, ",") {
		return input, ""
	}
	return input[:strings.Index(input, ",")], input[strings.Index(input, ",")+1:]
}

func findClosingIndex(input string) int {
	if input[0] != '[' {
		return 0
	}
	diffOpenClose := 1
	for i := 1; i < len(input); i++ {
		if input[i] == '[' {
			diffOpenClose++
		}
		if input[i] == ']' {
			diffOpenClose--
		}
		if diffOpenClose == 0 {
			return i
		}
	}
	return 0
}

func analyseElement(element string) int {
	if strings.Contains(element, ",") {
		return group
	}
	if element[0] != '[' {
		return number
	}
	if element[1] == ']' {
		return empty
	}
	return group
}

func checkInput(input []inputStrs) int {
	var sum int
	i := 1

	for _, pair := range input {
		if res, _ := strsAreOrdered(pair); res != unordered {
			sum += i
			//fmt.Printf("\nOrdered: %d - %s - %s\n", i, pair.left, pair.right)
		}
		i++
	}
	return sum
}

func sortPackages(pairs []inputStrs) []string {
	var finished bool
	res := make([]string, 0)
	for _, pair := range pairs {
		res = append(res, pair.left, pair.right)
	}
	res = append(res, dividerPackages[0], dividerPackages[1])
	for !finished {
		finished = true
		debug++
		//fmt.Printf("\n%d", debug)
		//if debug%100000 == 0 {
		//	fmt.Printf("%s", res)
		//}
		for i, p := range res {
			if i < len(res)-1 {
				if order, _ := strsAreOrdered(inputStrs{left: p, right: res[i+1]}); order == unordered {
					if debug%100000 == 0 {
						fmt.Printf("\n reodered: %d: %s - %s", i, p, res[i+1])
					}
					res[i] = res[i+1]
					res[i+1] = p
					finished = false
				}
			}
		}
	}
	return res
}

func findDividerPackages(input []string) int {
	var a, b int
	for i, p := range input {
		if p == dividerPackages[0] {
			a = i + 1
		}
		if p == dividerPackages[1] {
			b = i + 1
		}
	}
	return a * b
}
