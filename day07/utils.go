package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type lineType int

const (
	ls lineType = iota
	cdHome
	cdBack
	cd
	fileNS
	dir
)

type fileType struct {
	name string
	size int
}

type inputLineType struct {
	lineType        lineType
	currentPath     string
	file            fileType
	dir             string
	targetDirectory string
}

type directory struct {
	name        string
	path        string
	files       []fileType
	directories []*directory
}

func readIn(path string) ([]inputLineType, error) {
	input := make([]inputLineType, 0)
	currentDir := "/"
	file, err := os.Open(path)
	if err != nil {

		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		inputString := scanner.Text()
		iL := inputLineType{currentPath: currentDir}
		switch inputString[0] {
		case '$': //command
			if inputString[2] == 'l' {
				iL.lineType = ls
			} else {
				switch inputString[5] {
				case '/':
					iL.lineType = cdHome
					currentDir = "/"
				case '.':
					iL.lineType = cdBack
					strings.LastIndex(currentDir, "/")
					currentDir = currentDir[:strings.LastIndex(currentDir, "/")]
				default:
					iL.lineType = cd
					iL.targetDirectory = inputString[5:]
					currentDir += "/" + inputString[5:]
				}
			}
		case 'd': // directory
			iL.lineType = dir
			iL.dir = strings.Split(inputString, " ")[1]
			iL.currentPath = currentDir
		default: // file
			iL.lineType = fileNS
			strs := strings.Split(inputString, " ")
			size, err := strconv.ParseInt(strs[0], 10, 64)
			if err != nil {
				return input, err
			}
			iL.file.size = int(size)
			iL.file.name = strs[1]
			iL.currentPath = currentDir
		}
		input = append(input, iL)
	}
	return input, nil
}

func buildDirStructure(inputLines []inputLineType) directory {
	homeDir := directory{
		name: "/",
		path: "/",
	}
	currentDir := &homeDir
	for _, inputLine := range inputLines {
		switch inputLine.lineType {
		case ls:
			break
		case dir:
			insertDirIfNotExists(currentDir, inputLine.dir)
		case cd:
			insertDirIfNotExists(currentDir, inputLine.targetDirectory)
			currentDir = getDirByName(currentDir, inputLine.targetDirectory)
		case cdHome:
			currentDir = &homeDir
		case cdBack:
			currentDir = oneLevelBack(&homeDir, currentDir)
		case fileNS:
			insertFile(currentDir, inputLine.file)
		}

	}
	return homeDir
}

func insertDirIfNotExists(currentDir *directory, dirName string) {
	var separator string
	dirs := currentDir.directories
	exists := false
	for _, d := range dirs {
		if d.name == dirName {
			exists = true
		}
	}
	if !exists {
		if currentDir.path != "/" {
			separator = "/"
		}
		currentDir.directories = append(dirs, &directory{name: dirName, path: currentDir.path + separator + dirName})
	}
}

func getDirByName(currentDir *directory, dirName string) *directory {
	dirs := currentDir.directories
	for _, d := range dirs {
		if d.name == dirName {
			return d
		}
	}
	return nil
}

func oneLevelBack(home, currentDir *directory) *directory {
	dirNames := strings.Split(strings.TrimLeft(currentDir.path, "/"), "/")
	if len(dirNames) < 2 {
		return home
	}
	dir := home
	for i, dirName := range dirNames {
		dir = getDirByName(dir, dirName)
		if i == len(dirNames)-1-1 {
			return dir
		}

	}
	return currentDir
}

func insertFile(dir *directory, file fileType) {
	dir.files = append(dir.files, file)
}

func calcSingleDirSum(dir *directory) int {
	var sum int
	for _, file := range dir.files {
		sum += file.size
	}
	return sum
}

func calcRecursiveDirSum(dir *directory, ignoreAbove, ignoreBelow, sum int, results []int) (int, []int) {
	var subSum int
	newSum := sum
	res := results
	for _, subDir := range dir.directories {
		s, subRes := calcRecursiveDirSum(subDir, ignoreAbove, ignoreBelow, newSum, res)
		subSum += s
		res = subRes
	}
	s := subSum + calcSingleDirSum(dir)
	if s > ignoreBelow && s <= ignoreAbove {
		res = append(res, s)
	}
	return s, res

}

func showDir(directory *directory) {
	fmt.Printf("%s - %s", directory.name, directory.path)
	for _, f := range directory.files {
		fmt.Printf("\n%s - %d", f.name, f.size)
	}
	fmt.Println("\n++++++++++++")
	for _, d := range directory.directories {
		showDir(d)
	}

}
