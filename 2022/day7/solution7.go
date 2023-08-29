package main

import (
	"adventofcode2022/common"
	"fmt"
	"strconv"
	"strings"
)

type FileSystemEntry interface {
	Name() string
	Size() int
	Parent() FileSystemEntry
	SetParent(FileSystemEntry)
}

type Directory struct {
	FileSystemEntry
	name     string
	contents []FileSystemEntry
	parent   FileSystemEntry
}

type File struct {
	FileSystemEntry
	name   string
	size   int
	parent FileSystemEntry
}

func (f *File) Size() int {
	return f.size
}

func (f *File) Name() string {
	return f.name
}

func (f *File) Parent() FileSystemEntry {
	return f.parent
}

func (f *File) SetParent(fse FileSystemEntry) {
	f.parent = fse
}

func NewFile(name string, size int) *File {
	return &File{
		name: name,
		size: size,
	}
}

func (d *Directory) Size() int {
	var size int = 0

	for _, fse := range d.contents {
		size += fse.Size()
	}

	return size
}

func (d *Directory) Name() string {
	return d.name
}

func (d *Directory) Parent() FileSystemEntry {
	return d.parent
}

func (d *Directory) SetParent(fse FileSystemEntry) {
	d.parent = fse
}

func (d *Directory) Add(fse FileSystemEntry) {
	d.contents = append(d.contents, fse)
	fse.SetParent(d)
}

func NewDirectory(name string) *Directory {
	return &Directory{
		name: name,
	}
}

// TODO
// understand pointers and pass-by-value in Golang
// write some tests for building a directory tree here
// understand interfaces and structs better

func main() {
	var rootDirectory *Directory = NewDirectory("/")
	var currentDirectory *Directory = rootDirectory
	rootDirectory.SetParent(rootDirectory)
	var directories []*Directory
	directories = append(directories, rootDirectory)

	lineConsumer := func(line string) {
		switch line[0] {
		case '$':
			if line[2:4] == "cd" {
				currentDirectoryName := line[5:]

				switch currentDirectoryName {
				case "..":
					currentDirectory = currentDirectory.Parent().(*Directory)
				case "/":
					currentDirectory = rootDirectory
				default:
					newDirectory := NewDirectory(currentDirectoryName)
					currentDirectory.Add(newDirectory)
					currentDirectory = newDirectory
					directories = append(directories, currentDirectory)
				}
			}
		default: // read ls output
			var lsOutputParts = strings.Split(line, " ")
			if lsOutputParts[0] != "dir" {
				size, _ := strconv.Atoi(lsOutputParts[0])
				fileName := lsOutputParts[1]
				file := NewFile(fileName, size)
				currentDirectory.Add(file)
			}
		}
	}

	common.ReadFileLineByLine("input.txt", lineConsumer)

	// part 1
	fmt.Println("Riddle part 1")
	threshold := 100000
	sum := 0
	for _, dir := range directories {
		if sz := dir.Size(); sz <= threshold {
			sum += sz
		}
	}
	fmt.Println(sum)

	// part 2
	fmt.Println("Riddle part 2")
	totalDiskSpace := 70000000
	requiredSpace := 30000000
	currentDiskUtilization := rootDirectory.Size()
	currentFreeSpace := totalDiskSpace - currentDiskUtilization
	minimumCleanTarget := requiredSpace - currentFreeSpace
	currentRemovalTarget := rootDirectory

	for _, dir := range directories {
		if sz := dir.Size(); sz >= minimumCleanTarget && sz <= currentRemovalTarget.Size() {
			currentRemovalTarget = dir
		}
	}

	fmt.Println(currentRemovalTarget.Size())
}
