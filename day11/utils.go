package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const opAdd = "+"
const opMultiply = "*"
const opSquare = "square"
const opDoulbe = "double"

var smallestCommonMultiple int

type operartion struct {
	num int
	op  string
}

type monkeyType struct {
	num            int
	items          []int
	operartion     operartion
	testNum        int
	destMonkeys    []int
	numInspections int
}

type item struct {
	startValue int
	monkey     map[int][]int
}

func readIn(path string) ([]*monkeyType, error) {
	var n int64
	input := make([]*monkeyType, 0)
	file, err := os.Open(path)
	if err != nil {

		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 0
	monkey := new(monkeyType)
	dests := make([]int, 0)
	for scanner.Scan() {
		inputString := strings.TrimSpace(scanner.Text())
		strs := strings.Split(inputString, " ")
		switch i {
		case 0:
			n, err = strconv.ParseInt(strings.TrimRight(strs[1], ":"), 10, 64)
			if err != nil {
				return nil, err
			}
			monkey.num = int(n)
		case 1:
			items := make([]int, 0)
			for j := 2; j < len(strs); j++ {
				n, err = strconv.ParseInt(strings.TrimRight(strs[j], ","), 10, 64)
				if err != nil {
					return nil, err
				}
				items = append(items, int(n))
			}
			monkey.items = items
		case 2:
			if strs[5] == "old" {
				switch strs[4] {
				case opMultiply:
					monkey.operartion.op = opSquare
				case opAdd:
					monkey.operartion.op = opDoulbe
				}
			} else {
				monkey.operartion.op = strs[4]
				n, err = strconv.ParseInt(strs[5], 10, 64)
				if err != nil {
					return nil, err
				}
				monkey.operartion.num = int(n)
			}

		case 3:
			n, err = strconv.ParseInt(strs[3], 10, 64)
			if err != nil {
				return nil, err
			}
			monkey.testNum = int(n)
		case 4, 5:
			n, err = strconv.ParseInt(strs[5], 10, 64)
			if err != nil {
				return nil, err
			}
			dests = append(dests, int(n))
			if len(dests) == 2 {
				monkey.destMonkeys = dests
				input = append(input, monkey)
				dests = make([]int, 0)
				monkey = new(monkeyType)

			}
		case 6:
			i = -1
		}
		i++
	}
	smallestCommonMultiple = 1
	for _, m := range input {
		smallestCommonMultiple *= m.testNum
	}
	return input, nil
}

func inspectionRound(monkies []*monkeyType) []*monkeyType {
	var wLevel int
	for _, monkey := range monkies {
		for _, item := range monkey.items {
			monkey.numInspections++
			switch monkey.operartion.op {
			case opAdd:
				wLevel = item + monkey.operartion.num
			case opMultiply:
				wLevel = item * monkey.operartion.num
			case opDoulbe:
				wLevel = item + item
			case opSquare:
				wLevel = item * item
			}
			wLevel = wLevel / 3
			if wLevel%monkey.testNum == 0 {
				monkies[monkey.destMonkeys[0]].items = append(monkies[monkey.destMonkeys[0]].items, wLevel)
			} else {
				monkies[monkey.destMonkeys[1]].items = append(monkies[monkey.destMonkeys[1]].items, wLevel)
			}
		}
		monkey.items = make([]int, 0)
	}
	return monkies
}

func getMonkeyBusiness(monkies []*monkeyType) int {
	sort.Slice(monkies, func(i, j int) bool {
		return monkies[i].numInspections > monkies[j].numInspections
	})
	return monkies[0].numInspections * monkies[1].numInspections
}

func removeItem(items []*item, item *item) []*item {

	for i, v := range items {
		if v == item {
			fmt.Printf("\nremove: %s", v)
			return append(items[:i], items[i+1:]...)
		}
	}
	return items
}

func inspectionRoundPart2(monkies []*monkeyType) []*monkeyType {
	var wLevel int
	for _, monkey := range monkies {
		for _, item := range monkey.items {
			monkey.numInspections++
			switch monkey.operartion.op {
			case opAdd:
				wLevel = item + monkey.operartion.num
			case opMultiply:
				wLevel = item * monkey.operartion.num
			case opDoulbe:
				wLevel = item + item
			case opSquare:
				wLevel = item * item
			}
			wLevel = wLevel % smallestCommonMultiple
			if wLevel%monkey.testNum == 0 {
				monkies[monkey.destMonkeys[0]].items = append(monkies[monkey.destMonkeys[0]].items, wLevel)
			} else {
				monkies[monkey.destMonkeys[1]].items = append(monkies[monkey.destMonkeys[1]].items, wLevel)
			}
		}
		monkey.items = make([]int, 0)
	}
	return monkies
}
