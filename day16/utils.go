package main

import (
	"bufio"
	"golang.org/x/exp/slices"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var debug int

type cave struct {
	name     string
	flowRate int
	ways     []string
	visited  int
}

type cave2 struct {
	name     string
	flowRate int
	ways     map[string]int
	visited  int
}

func readIn(path string) (map[string]cave, error) {
	caves := make(map[string]cave)
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputString := strings.TrimSpace(scanner.Text())
		cave := strToCave(inputString)
		caves[cave.name] = cave
	}
	return caves, nil
}

func strToCave(s string) cave {
	strs := strings.Split(s, " ")
	name := strs[1]
	flowRate, _ := strconv.ParseInt(strs[4][strings.Index(strs[4], "=")+1:len(strs[4])-1], 10, 64)
	ways := make([]string, 0)
	for i := 9; i < len(strs); i++ {
		str := strings.TrimRight(strs[i], ",")
		ways = append(ways, str)
	}
	return cave{name: name, flowRate: int(flowRate), ways: ways}
}

func openingSinceLastVisit(path, dest string) bool {
	if len(path) < 2 {
		return false
	}
	for i := len(path) - 2; i > 0; i = i - 2 {
		if path[i:i+2] == dest {
			return false
		}
		if path[i:i+2] == openingStr {
			return true
		}
	}
	return false
}

func isOpen(path, cave string) bool {
	return strings.Contains(path, cave+openingStr)
}

func wayArrayToMap(caves map[string]cave, from []string, dest string, cost int, wayMap map[string]int) map[string]int {
	//if dest == start {
	//	wayMap[dest] = cost
	//	return wayMap
	//}
	if caves[dest].flowRate != 0 {
		wayMap[dest] = cost
		return wayMap
	}
	for _, nextWay := range caves[dest].ways {
		if slices.Contains(from, nextWay) {
			continue
		}
		from = append(from, dest)
		wayMap = wayArrayToMap(caves, from, nextWay, cost+1, wayMap)
	}
	return wayMap
}

func rebuildMap(caves map[string]cave) map[string]cave2 {
	c2s := make(map[string]cave2)
	for name, c := range caves {
		if c.flowRate == 0 && c.name != start {
			continue
		}
		ways2 := make(map[string]int)
		for _, oldway := range c.ways {
			from := []string{name}
			ways2 = wayArrayToMap(caves, from, oldway, 1, ways2)
			c2 := cave2{
				name:     c.name,
				flowRate: c.flowRate,
				ways:     ways2,
				visited:  c.visited,
			}
			c2s[name] = c2
		}

	}
	return c2s
}

func findMax2(caves map[string]cave2, remaining int, start string, rate, sum int, visited []string, path string) (resSum int) {
	var newSumGoingDirectly, newSumOpeningFirst, partSum int

	cave := caves[start]
	visited = append(visited, cave.name)
	//fmt.Printf("\nHistory: %s", history)
	if remaining <= 0 {
		return sum
	}
	path += cave.name
	newSumOpeningFirst = sum + rate
	newSumGoingDirectly = sum + rate
	//Max when going to other caves directly
	for dest, cost := range cave.ways {
		if cost > remaining {
			continue
		}
		if slices.Contains(visited, dest) && !openingSinceLastVisit(path, dest) {
			//fmt.Printf("\nskipping: %s because of path %s", dest, path)
			continue
		}
		partSum = findMax2(caves, remaining-cost, dest, rate, sum+(cost*rate), visited, path)
		if partSum > newSumGoingDirectly {
			newSumGoingDirectly = partSum
		}
	}
	//Max when opening valve first
	if !isOpen(path, cave.name) && cave.flowRate > 0 && remaining > 1 {
		//fmt.Printf("\nOpening: %s on path %s", cave.name, path)
		pathOpening := path + openingStr

		for dest, cost := range cave.ways {
			if cost > remaining-1 {
				continue
			}

			partSum = findMax2(caves, remaining-1-cost, dest, rate+cave.flowRate, sum+rate+cost*(rate+cave.flowRate), visited, pathOpening)
			if partSum > newSumOpeningFirst {
				newSumOpeningFirst = partSum

			}
		}

		if newSumOpeningFirst > newSumGoingDirectly {
			return newSumOpeningFirst
		}
	}
	return newSumGoingDirectly
}

func findMaxWithHelp(caves map[string]cave2, remaining int, start1, start2 string, sum, idle1, idle2 int, path1, path2 string) (resSum int) {
	if allValvesOpen(caves, path1+path2) || remaining == 0 {
		return sum
	}
	switch {
	case idle1 == 0:
		switch {
		case idle2 == 0: // 1s2w
			return findMaxWithHelp(caves, remaining-1, start1, start2, sum, idle1-1, idle2-1, path1, path2)

		case idle2 == -1: // 1o2s
			return findMaxOneIdle(caves, remaining, start2, start1, sum, idle1, path1, path2)
		case idle2 > 0: // 1s2w
			return findMaxWithHelp(caves, remaining-1, start1, start2, sum, idle1-1, idle2-1, path1, path2)
		}

	case idle1 == -1: // is open ->search new
		switch {
		case idle2 == 0: //1s2o
			return findMaxOneIdle(caves, remaining, start1, start2, sum, idle2, path2, path1)
		case idle2 == -1: // both searching
			return findMaxBothSearching(caves, remaining, start1, start2, sum, idle1, idle2, path1, path2)
		case idle2 > 0: // 1s2w
			return findMaxOneIdle(caves, remaining, start1, start2, sum, idle2, path1, path2)
		}
	case idle1 > 0: // waiting
		switch {
		case idle2 == 0: // 1s2w
			return findMaxWithHelp(caves, remaining-1, start1, start2, sum, idle1-1, idle2-1, path1, path2)

		case idle2 == -1: // 1w2s
			return findMaxOneIdle(caves, remaining, start2, start1, sum, idle1, path2, path1)
		case idle2 > 0: //both waiting
			minIdle := getMinInt(idle1, idle2)
			minIdle = getMinInt(minIdle, remaining)
			return findMaxWithHelp(caves, remaining-minIdle, start1, start2, sum, idle1-minIdle, idle2-minIdle, path1, path2)
		}
	}
	log.Fatalf("\nUNEXPECTED CASE (remaining %d) %d | %d\n", remaining, idle1, idle2)
	return
}

func getMaxInt(a1, a2 int) int {
	return int(math.Max(float64(a1), float64(a2)))
}

func getMinInt(a1, a2 int) int {
	return int(math.Min(float64(a1), float64(a2)))
}

func rebuildMapToIncludeOnlyValves(caves map[string]cave2) map[string]cave2 {
	newCaves := make(map[string]cave2)
	for caveName, cave := range caves {
		if cave.flowRate == 0 && caveName != start {
			continue
		}
		waysToValves := make(map[string]int)
		for namePotentialWay, potentialWay := range caves {
			if namePotentialWay == caveName {
				continue
			}
			if potentialWay.flowRate == 0 {
				continue
			}
			waysToValves[namePotentialWay] = findShortestWay(caves, caveName, namePotentialWay, "", 0)
		}

		newCave := cave2{
			name:     cave.name,
			flowRate: cave.flowRate,
			ways:     waysToValves,
			visited:  0,
		}
		newCaves[caveName] = newCave
	}
	return newCaves
}

func findShortestWay(caves map[string]cave2, start, end, path string, costsToStart int) int {
	minCosts := -1
	path += start
	if start == end {
		return costsToStart
	}

	for w, cost := range caves[start].ways {
		if strings.Contains(path, w) {
			continue
		}
		if w == end {
			return costsToStart + cost
		}
		partCosts := findShortestWay(caves, w, end, path, costsToStart+cost)
		if partCosts == -1 {
			continue
		}
		if minCosts == -1 || partCosts < minCosts {
			minCosts = partCosts
		}
	}
	return minCosts

}

func findMaxOneIdle(caves map[string]cave2, remaining int, start, startIdle string, sum, idle int, path, pathIdle string) int {
	newSum := sum
	found := false
	for dest, cost := range caves[start].ways {
		if cost >= remaining {
			continue
		}
		if isOpen(path+pathIdle, dest) {
			continue
		}
		if dest == startIdle {
			continue
		}
		found = true
		partSum := findMaxWithHelp(caves, remaining-1, dest, startIdle, sum+(remaining-cost-1)*caves[dest].flowRate, cost-1, idle-1, path+dest+openingStr, pathIdle)
		if partSum > newSum {
			newSum = partSum
		}
	}
	if !found && remaining > idle && !allValvesOpen(caves, path+pathIdle) {
		partSum := findMaxWithHelp(caves, remaining-idle-1, start, startIdle, sum, remaining+100, 0, path, pathIdle)
		if partSum > newSum {
			newSum = partSum
		}
	}
	return newSum
}

func findMaxBothSearching(caves map[string]cave2, remaining int, start1, start2 string, sum, idle1, idle2 int, path1, path2 string) int {
	cave1 := caves[start1]
	cave2 := caves[start2]
	newSum := sum
	for dest1, cost1 := range cave1.ways {
		if cost1 >= remaining {
			continue
		}
		if isOpen(path1+path2, dest1) {
			continue
		}
		for dest2, cost2 := range cave2.ways {
			if cost2 >= remaining {
				continue
			}
			if dest1 == dest2 {
				continue
			}
			if isOpen(path1+path2, dest2) {
				continue
			}
			minCost := getMinInt(cost1, cost2)
			s := sum + (remaining-cost1-1)*caves[dest1].flowRate + (remaining-cost2-1)*caves[dest2].flowRate
			partSum := findMaxWithHelp(caves, remaining-minCost, dest1, dest2, s, cost1-minCost, cost2-minCost, path1+dest1+openingStr, path2+dest2+openingStr)
			if partSum > newSum {
				newSum = partSum
			}
		}
	}
	return newSum
}

func allValvesOpen(caves map[string]cave2, path string) bool {
	for cName, c := range caves {
		if c.flowRate == 0 {
			continue
		}
		if !isOpen(path, cName) {
			return false
		}
	}
	return true
}
