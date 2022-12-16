package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	x int
	y int
}

type line struct {
	start coordinate
	end   coordinate
}

func (l line) includes(c coordinate) bool {
	xDistanceFromStart := c.x - l.start.x
	return c.y == l.start.y+xDistanceFromStart
}

func (l line) getCoordinatesWithCap(cap int) []coordinate {
	res := make([]coordinate, 0)
	if l.runsTopDownToRight() {
		if l.start.x < l.end.x {
			for i := 0; i <= l.end.x-l.start.x; i++ {
				if l.start.x+i < 0 || l.start.y+i < 0 {
					continue
				}
				if l.start.x+i > cap || l.start.y+i > cap {
					break
				}
				res = append(res, coordinate{x: l.start.x + i, y: l.start.y + i})
			}
		} else {
			for i := 0; i <= l.start.x-l.end.x; i++ {
				if l.end.x+i < 0 || l.end.y+i < 0 {
					continue
				}
				if l.end.x+i > cap || l.end.y+i > cap {
					break
				}
				res = append(res, coordinate{x: l.end.x + i, y: l.end.y + i})
			}
		}
	} else {
		if l.start.x < l.end.x {
			for i := 0; i <= l.end.x-l.start.x; i++ {
				if l.start.x+i < 0 || l.start.y-i > cap {
					continue
				}
				if l.start.x+i > cap || l.start.y-i < 0 {
					break
				}
				res = append(res, coordinate{x: l.start.x + i, y: l.start.y - i})
			}
		} else {
			for i := 0; i <= l.start.x-l.end.x; i++ {
				if l.end.x+i < 0 || l.end.y-i > cap {
					continue
				}
				if l.end.x+i > cap || l.end.y-i < 0 {
					break
				}
				res = append(res, coordinate{x: l.end.x + i, y: l.end.y - i})
			}
		}
	}
	return res
}

func (l line) overlaps(l2 line) *line {
	if l.runsTopDownToRight() != l2.runsTopDownToRight() {
		//diagonal
		return nil
	}
	var minX2, minX2Y, maxX, maxXY int
	minX := getMinInt(l.start.x, l.end.x)

	if l.start.x > l.end.x {
		maxX = l.start.x
		maxXY = l.start.y
	} else {
		maxX = l.end.x
		maxXY = l.end.y
	}

	if l2.start.x < l2.end.x {
		minX2 = l2.start.x
		minX2Y = l2.start.y
	} else {
		minX2 = l2.end.x
		minX2Y = l2.start.y
	}
	if minX < minX2 {
		if maxX > minX2 {
			return &line{start: coordinate{x: minX2, y: minX2Y}, end: coordinate{x: maxX, y: maxXY}}
		}
	}
	return nil
}

func (l line) runsTopDownToRight() bool {
	return (l.start.x < l.end.x && l.start.y < l.end.y) || (l.start.x > l.end.x && l.start.y > l.end.y)
}

func (l line) crossesAt(l2 line) *coordinate {
	if l.runsTopDownToRight() == l2.runsTopDownToRight() {
		//parallel
		return nil
	}
	lTopDownToRight := l
	lTopDownToLeft := l2
	if l2.runsTopDownToRight() {
		lTopDownToRight = l2
		lTopDownToLeft = l
	}
	startR := lTopDownToRight.start
	endR := lTopDownToRight.end
	if startR.y > endR.y {
		startR = lTopDownToRight.end
	}
	startL := lTopDownToLeft.start
	endL := lTopDownToLeft.end
	if startL.y < endL.y {
		startL = lTopDownToLeft.start
	}
	d := startL.y - startR.y + startL.x - startR.x
	if d%2 != 0 {
		return nil
	}
	if startL.x+d > getMaxInt(lTopDownToLeft.start.x, lTopDownToLeft.end.x) {
		return nil
	}
	if startL.x+d < getMinInt(lTopDownToLeft.start.x, lTopDownToLeft.end.x) {
		return nil
	}
	if startL.y+d > getMaxInt(lTopDownToLeft.start.y, lTopDownToLeft.end.y) {
		return nil
	}
	if startL.y+d < getMinInt(lTopDownToLeft.start.y, lTopDownToLeft.end.y) {
		return nil
	}
	return &coordinate{x: startL.x + d, y: startL.y - d}

}

func (c coordinate) getDistance(c2 coordinate) int {
	return int(math.Abs(float64(c.x-c2.x)) + math.Abs(float64(c.y-c2.y)))
}
func (c coordinate) getDistanceX(c2 coordinate) int {
	return int(math.Abs(float64(c.x - c2.x)))
}
func (c coordinate) getDistanceY(c2 coordinate) int {
	return int(math.Abs(float64(c.y - c2.y)))
}

type outerRing struct {
	topToRight    line
	rightToBottom line
	bottomToLeft  line
	leftToTop     line
}

type sensorBeacon struct {
	sensor coordinate
	beacon coordinate
}

func (sB sensorBeacon) getXDistance() int {
	return int(math.Abs(float64(sB.sensor.x - sB.beacon.x)))
}
func (sB sensorBeacon) getYDistance() int {
	return int(math.Abs(float64(sB.sensor.y - sB.beacon.y)))
}

func (sB sensorBeacon) getSmallerX() int {
	return int(math.Min(float64(sB.sensor.x), float64(sB.beacon.x)))
}

func (sB sensorBeacon) getSmallerY() int {
	return int(math.Min(float64(sB.sensor.y), float64(sB.beacon.y)))
}

func (sB sensorBeacon) getRadius() int {
	return sB.getXDistance() + sB.getYDistance()
}

func (sB sensorBeacon) getOuterRingWithCap(cap int) outerRing {
	var res outerRing

	lXRight := sB.sensor.x + sB.getRadius() + 1
	lXLeft := sB.sensor.x - sB.getRadius() - 1
	lyTop := sB.sensor.y + sB.getRadius() + 1
	lyBottom := sB.sensor.y - sB.getRadius() - 1

	top := considerCap(coordinate{x: sB.sensor.x, y: lyTop}, 0, cap)
	bottom := considerCap(coordinate{x: sB.sensor.x, y: lyBottom}, 0, cap)
	right := considerCap(coordinate{x: lXRight, y: sB.sensor.y}, 0, cap)
	left := considerCap(coordinate{x: lXLeft, y: sB.sensor.y}, 0, cap)

	res.topToRight = line{start: top, end: right}
	res.rightToBottom = line{start: right, end: bottom}
	res.bottomToLeft = line{start: bottom, end: left}
	res.leftToTop = line{start: left, end: top}

	return res
}

func considerCap(c coordinate, minCap, maxCap int) coordinate {
	x := c.x
	y := c.y
	if x < minCap {
		x = minCap
	}
	if x > maxCap {
		x = maxCap
	}
	if y < minCap {
		y = minCap
	}
	if y > maxCap {
		y = maxCap
	}
	return coordinate{x: x, y: y}
}

const (
	empty = iota
	sensor
	beacon
	noBeacon
)

func readIn(path string) ([]sensorBeacon, error) {
	var cS, cB coordinate
	input := make([]sensorBeacon, 0)
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputString := strings.TrimSpace(scanner.Text())
		cS, cB = strToCoordinate(inputString)
		input = append(input, sensorBeacon{sensor: cS, beacon: cB})
	}
	return input, nil
}

func strToCoordinate(s string) (sensor coordinate, beacon coordinate) {
	strs := strings.Split(s, " ")

	xS, _ := strconv.ParseInt(strs[2][2:len(strs[2])-1], 10, 64)
	yS, _ := strconv.ParseInt(strs[3][2:len(strs[3])-1], 10, 64)
	xB, _ := strconv.ParseInt(strs[8][2:len(strs[8])-1], 10, 64)
	yB, _ := strconv.ParseInt(strs[9][2:len(strs[9])], 10, 64)
	return coordinate{x: int(xS), y: int(yS)}, coordinate{x: int(xB), y: int(yB)}
}

func buildGrid(sensorBeacons []sensorBeacon) map[coordinate]int {
	grid := make(map[coordinate]int)
	for _, sB := range sensorBeacons {
		distance := sB.getYDistance() + sB.getXDistance()
		for i := 0; i <= distance; i++ {
			for j := 0; j <= distance-i; j++ {
				grid[coordinate{x: sB.sensor.x + j, y: sB.sensor.y + i}] = noBeacon
				grid[coordinate{x: sB.sensor.x + j, y: sB.sensor.y - i}] = noBeacon
				grid[coordinate{x: sB.sensor.x - j, y: sB.sensor.y + i}] = noBeacon
				grid[coordinate{x: sB.sensor.x - j, y: sB.sensor.y - i}] = noBeacon
			}
		}
		grid[sB.beacon] = beacon
		grid[sB.sensor] = sensor
	}
	return grid
}

func visualizeGrid(grid map[coordinate]int) {
	var minX, minY, maxX, maxY int
	for c, _ := range grid {
		if c.x < minX {
			minX = c.x
		}
		if c.x > maxX {
			maxX = c.x
		}
		if c.y < minY {
			minY = c.y
		}
		if c.y > maxY {
			maxY = c.y
		}
	}

	for rowIndex := minY; rowIndex <= maxY-minY; rowIndex++ {
		fmt.Print("\n")
		for i := minX; i <= maxX-minX; i++ {
			if grid[coordinate{x: i, y: rowIndex}] == empty {
				fmt.Print(" .")
				continue
			}
			if grid[coordinate{x: i, y: rowIndex}] == noBeacon {
				//fmt.Print("X")
				fmt.Printf("%2d", i)
				continue
			}
			if grid[coordinate{x: i, y: rowIndex}] == sensor {
				if rowIndex >= 20 {
					fmt.Printf(" R")
					continue
				}
				if rowIndex == 11 {
					fmt.Printf(" U")
					continue
				}
				fmt.Print(" S")
				continue
			}
			if grid[coordinate{x: i, y: rowIndex}] == beacon {
				fmt.Print(" B")
				continue
			}
			fmt.Print(" .")
		}

	}
	fmt.Printf("\n")

}

func visualizeRow(grid map[coordinate]int, rowIndex, width int) {

	for i := 0; i <= width; i++ {
		if grid[coordinate{x: i, y: rowIndex}] == empty {
			fmt.Print(".")
			continue
		}
		if grid[coordinate{x: i, y: rowIndex}] == noBeacon {
			fmt.Print("X")
			continue
		}
		if grid[coordinate{x: i, y: rowIndex}] == sensor {
			if rowIndex >= 20 {
				fmt.Printf("R")
				continue
			}
			fmt.Print("S")
			continue
		}
		if grid[coordinate{x: i, y: rowIndex}] == beacon {
			fmt.Print("B")
			continue
		}
		fmt.Print(".")
	}
	fmt.Print("\n")
}

func countRow(grid map[coordinate]int, rowIndex int) int {
	var sum int
	for c, v := range grid {
		if c.y == rowIndex && (v == noBeacon || v == sensor) {
			sum++
		}
	}
	return sum
}

func getCanidadates(sensorBeacons []sensorBeacon, cap int) []coordinate {
	candidates := make([]coordinate, 0)
	for _, sB1 := range sensorBeacons {
		outerRing1 := sB1.getOuterRingWithCap(cap)
		for _, sB2 := range sensorBeacons {
			if sB1.sensor.getDistanceX(sB2.sensor) == 0 {
				continue
			}
			outerRing2 := sB2.getOuterRingWithCap(cap)
			if c := outerRing1.rightToBottom.crossesAt(outerRing2.bottomToLeft); c != nil {
				candidates = append(candidates, *c)
			}
			if c := outerRing1.leftToTop.crossesAt(outerRing2.bottomToLeft); c != nil {
				candidates = append(candidates, *c)
			}
			if c := outerRing1.rightToBottom.crossesAt(outerRing2.topToRight); c != nil {
				candidates = append(candidates, *c)
			}
			if c := outerRing1.leftToTop.crossesAt(outerRing2.topToRight); c != nil {
				candidates = append(candidates, *c)
			}
			//xxxxxxParallel
			l1 := outerRing1.rightToBottom.overlaps(outerRing2.leftToTop)
			l2 := outerRing1.bottomToLeft.overlaps(outerRing2.topToRight)
			l3 := outerRing1.leftToTop.overlaps(outerRing2.rightToBottom)
			l4 := outerRing1.topToRight.overlaps(outerRing2.bottomToLeft)

			if l1 != nil {
				candidates = append(candidates, l1.getCoordinatesWithCap(cap)...)
			}
			if l2 != nil {
				candidates = append(candidates, l2.getCoordinatesWithCap(cap)...)
			}
			if l3 != nil {
				candidates = append(candidates, l3.getCoordinatesWithCap(cap)...)
			}
			if l4 != nil {
				candidates = append(candidates, l4.getCoordinatesWithCap(cap)...)
			}

		}

	}

	return candidates
}

func buildRow(sensorBeacons []sensorBeacon, rowIndex int) map[coordinate]int {
	grid := make(map[coordinate]int)
	for _, sB := range sensorBeacons {
		distance := sB.getXDistance() + sB.getYDistance()
		yOffset := int(math.Abs(float64(sB.sensor.y - rowIndex)))
		if distance < yOffset {
			continue
		}
		for i := 0; i <= distance-yOffset; i++ {
			grid[coordinate{x: sB.sensor.x - i, y: rowIndex}] = noBeacon
			grid[coordinate{x: sB.sensor.x + i, y: rowIndex}] = noBeacon

		}
		if sB.beacon.y == rowIndex {
			grid[sB.beacon] = beacon
		}
		if sB.sensor.y == rowIndex {
			grid[sB.sensor] = sensor
		}

	}
	return grid
}

func getMinInt(a1, a2 int) int {
	return int(math.Min(float64(a1), float64(a2)))
}
func getMaxInt(a1, a2 int) int {
	return int(math.Max(float64(a1), float64(a2)))
}
