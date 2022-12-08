package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type File struct {
	Name string
	Size int
}

type Directory struct {
	Name   string
	Parent *Directory

	Files       map[string]*File
	Directories map[string]*Directory

	cachedSize    int
	hasCachedSize bool
}

func (d *Directory) Size() int {
	if d.hasCachedSize {
		return d.cachedSize
	}

	size := 0

	for _, file := range d.Files {
		size += file.Size
	}

	for _, sd := range d.Directories {
		size += sd.Size()
	}

	d.hasCachedSize = true
	d.cachedSize = size

	return size
}

func (d *Directory) TotalSize(lim int) int {
	res := 0

	if d.Size() <= lim {
		res += d.Size()
	}

	for _, sd := range d.Directories {
		res += sd.TotalSize(lim)
	}

	return res
}

func (d *Directory) IsCloser(size int, od *Directory) bool {
	if od == nil {
		return true
	}

	currentDiff := size - d.Size()
	diff := size - od.Size()

	if currentDiff < 0 && diff < 0 && diff > currentDiff {
		return false
	} else if currentDiff > 0 && diff < currentDiff {
		return false
	}

	return true
}

func (d *Directory) FindClosest(size int, currentCandidate *Directory) *Directory {
	if d.IsCloser(size, currentCandidate) {
		currentCandidate = d
	}

	for _, d := range d.Directories {
		currentCandidate = d.FindClosest(size, currentCandidate)
	}

	return currentCandidate
}

func main() {
	input := loadInput("input.txt")

	log.Println(part1(input))
	log.Println(part2(input))
}

func part1(input *Directory) int {
	return input.TotalSize(100000)
}

func part2(input *Directory) int {
	total := 70000000
	target := 30000000
	used := input.Size()

	unused := total - used
	toFree := target - unused

	d := input.FindClosest(toFree, nil)

	return d.Size()
}

func loadInput(path string) *Directory {
	file, _ := os.Open(path)
	defer file.Close()

	var currentDirectory *Directory

	rootDirectory := &Directory{
		Name:        "/",
		Directories: make(map[string]*Directory),
		Files:       make(map[string]*File),
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "$ cd /" {
			currentDirectory = rootDirectory
		} else if line == "$ cd .." {
			currentDirectory = currentDirectory.Parent
		} else if line == "$ ls" {
			continue
		} else if strings.Contains(line, "$ cd ") {
			name := line[5:]
			currentDirectory = currentDirectory.Directories[name]
		} else if strings.Contains(line, "dir ") {
			name := line[4:]
			if _, ok := currentDirectory.Directories[name]; !ok {
				currentDirectory.Directories[name] = &Directory{
					Name:   name,
					Parent: currentDirectory,

					Files:       make(map[string]*File),
					Directories: make(map[string]*Directory),
				}
			}
		} else {
			parts := strings.Split(line, " ")
			size, _ := strconv.Atoi(parts[0])

			currentDirectory.Files[parts[1]] = &File{
				Name: parts[1],
				Size: size,
			}
		}
	}

	return rootDirectory
}
