package main

const rowIndexExample = 10
const rowIndex = 2000000

func part1(filePath string) (int, error) {
	sensorBeacons, err := readIn(filePath)
	if err != nil {
		return 0, err
	}
	grid := buildRow(sensorBeacons, rowIndex)

	return countRow(grid, rowIndex), nil
}

func part1Example(filePath string) (int, error) {
	sensorBeacons, err := readIn(filePath)
	if err != nil {
		return 0, err
	}

	grid := buildGrid(sensorBeacons)
	visualizeGrid(grid)
	return countRow(grid, rowIndexExample), nil
}
