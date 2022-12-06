package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func readIn(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {

		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		inputString := scanner.Text()
		return inputString, nil
	}
	return "", nil
}

func afterUnique(input string, numberOfUniqueChar int) int {
	if len(input) < 4 {
		return 0
	}

	for i := 0; i < len(input)-numberOfUniqueChar; i++ {
		unique := true
		for j := 0; j < numberOfUniqueChar; j++ {
			if strings.Contains(input[i+j+1:i+numberOfUniqueChar], string(input[i+j])) {
				unique = false
				break
			}
		}
		if unique {
			return i + numberOfUniqueChar
		}
	}
	return 0
}
