package main

import "fmt"

const searchSpace = 20
const searchSpace2 = 4000000

func part2(filePath string) (int, error) {
	candidates := make([]coordinate, 0)
	sensorBeacons, err := readIn(filePath)
	if err != nil {
		return 0, err
	}

	candidates = getCanidadates(sensorBeacons, searchSpace2)
candiateLoop:
	for _, c := range candidates {
		for _, sB := range sensorBeacons {
			if sB.sensor.getDistance(c) <= sB.getXDistance()+sB.getYDistance() {
				continue candiateLoop
			}
		}
		fmt.Printf("Empty coordinate: %d - %d\n", c.x, c.y)
		return c.x*searchSpace2 + c.y, nil
	}

	return -1, nil
}
