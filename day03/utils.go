package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

type Compartment struct {
	left  string
	right string
}

type ElveGroup struct {
	first  string
	second string
	third  string
}

func (c Compartment) GetIdentical() rune {
	for _, s := range c.left {
		if strings.Contains(c.right, fmt.Sprintf("%c", s)) {
			return rune(s)
		}
	}
	return 0
}

func (eG ElveGroup) GetIdentical() rune {
	var common12 string
	for _, s := range eG.first {
		if strings.Contains(eG.second, fmt.Sprintf("%c", s)) {
			common12 += fmt.Sprintf("%c", s)
		}
	}
	for _, s := range eG.third {
		if strings.Contains(common12, fmt.Sprintf("%c", s)) {
			return rune(s)
		}
	}

	return 0
}

func readIn(path string) ([]Compartment, error) {
	file, err := os.Open(path)
	if err != nil {

		log.Fatal(err)
	}
	defer file.Close()
	compartments := make([]Compartment, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputString := strings.TrimSpace(scanner.Text())
		c := Compartment{
			left:  inputString[:(len(inputString) / 2)],
			right: inputString[len(inputString)/2:],
		}
		compartments = append(compartments, c)
	}
	return compartments, nil
}

func readInPart2(path string) ([]ElveGroup, error) {
	file, err := os.Open(path)
	if err != nil {

		log.Fatal(err)
	}
	defer file.Close()
	elveGroups := make([]ElveGroup, 0)
	scanner := bufio.NewScanner(file)
	var i int
	eg := new(ElveGroup)
	for scanner.Scan() {
		i++
		inputString := strings.TrimSpace(scanner.Text())

		switch i {
		case 1:
			eg.first = inputString
		case 2:
			eg.second = inputString
		case 3:
			eg.third = inputString
			elveGroups = append(elveGroups, *eg)
			eg = new(ElveGroup)
			i = 0
		}
	}
	return elveGroups, nil
}

func getCharacterValue(r rune) int {
	if unicode.IsUpper(r) {
		return int(r) - 64 + 26
	}
	return int(r) - 96
}

func getScore(compartments []Compartment) int {
	var sum int
	for _, c := range compartments {
		sum += getCharacterValue(c.GetIdentical())
	}
	return sum
}

func getScorePart2(elveGroups []ElveGroup) int {
	var sum int
	for _, eG := range elveGroups {
		sum += getCharacterValue(eG.GetIdentical())
	}
	return sum
}
