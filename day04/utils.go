package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type ElveGroup struct {
	first  []int
	second []int
}

func readIn(path string) ([]ElveGroup, error) {
	file, err := os.Open(path)
	if err != nil {

		log.Fatal(err)
	}
	defer file.Close()
	elvegroups := make([]ElveGroup, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		eg := new(ElveGroup)
		inputString := strings.TrimSpace(scanner.Text())
		strs := strings.Split(inputString, ",")
		first, err := getIds(strs[0])
		if err != nil {
			return elvegroups, err
		}
		second, err := getIds(strs[1])
		if err != nil {
			return elvegroups, err
		}
		eg.first = first
		eg.second = second
		elvegroups = append(elvegroups, *eg)
	}
	return elvegroups, nil
}

func getIds(input string) ([]int, error) {
	strs := strings.Split(input, "-")
	ids := make([]int, 0)
	begin, err := strconv.ParseInt(strs[0], 10, 64)
	if err != nil {
		return ids, err
	}
	end, err := strconv.ParseInt(strs[1], 10, 64)
	if err != nil {
		return ids, err
	}
	for i := begin; i <= end; i++ {
		ids = append(ids, int(i))
	}
	return ids, nil
}

func fullyContains(first, second []int) bool {
	return first[0] <= second[0] && first[len(first)-1] >= second[len(second)-1]
}

func partiallyContains(first, second []int) bool {
	return first[0] <= second[len(second)-1] && first[len(first)-1] >= second[0]
}

func countFullyContains(elvegroups []ElveGroup) int {
	var num int
	for _, eG := range elvegroups {
		if fullyContains(eG.first, eG.second) || fullyContains(eG.second, eG.first) {
			num++
		}
	}
	return num
}

func countPartiallyContains(elvegroups []ElveGroup) int {
	var num int
	for _, eG := range elvegroups {
		if partiallyContains(eG.first, eG.second) || partiallyContains(eG.second, eG.first) {
			num++
		}
	}
	return num
}
