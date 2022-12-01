package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func readIn(path string) (map[int]int, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	valueMap := make(map[int]int, 0)
	scanner := bufio.NewScanner(file)

	var (
		elveSum     int
		elveNumber  int
		inputString string
	)
	for scanner.Scan() {
		inputString := strings.TrimSpace(scanner.Text())
		if len(inputString) == 0 {
			valueMap[elveNumber] = elveSum
			elveSum = 0
			elveNumber++
		} else {
			row, err := strconv.Atoi(inputString)
			if err != nil {
				return valueMap, err
			}
			elveSum += int(row)
		}
	}
	if len(inputString) != 0 {
		valueMap[elveNumber] = elveSum

	}
	return valueMap, nil
}

func getMax(input map[int]int) int {
	var max int
	for _, val := range input {
		if val > max {
			max = val
		}
	}
	return max
}

func sortArray(input map[int]int) map[int]int {
	found := true
	for found {

		found = false
		for i, val := range input {
			if int(len(input)) > i+1 {
				if val > input[i+1] {
					input[i] = input[i+1]
					input[i+1] = val
					found = true
				}
			}

		}
	}

	return input
}

func getTopThreeSum(input map[int]int) int {
	maxMap := make(map[int]int, 3)
	var i int
	for i = 0; i < 3; i++ {
		maxMap[i] = 0
	}
	for _, val := range input {
		if val > maxMap[0] {
			maxMap[0] = val
			maxMap = sortArray(maxMap)
		}
	}
	var sum int
	for _, val := range maxMap {
		sum += val
	}
	return sum
}
