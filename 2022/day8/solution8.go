package main

import (
	"adventofcode2022/common"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var grid [][]int

	var lineConsumer = func(line string) {
		var gridLine []int
		gridCharacters := strings.Split(line, "")

		for _, s := range gridCharacters {
			i, _ := strconv.Atoi(s)
			gridLine = append(gridLine, i)
		}

		grid = append(grid, gridLine)
	}

	common.ReadFileLineByLine("input.txt", lineConsumer)
	edgeTreeCount := 2*len(grid[0]) + 2*len(grid) - 4

	// calculate visiblity for each tree from all angles
	visibleInnerTrees := 0
	for y := 1; y <= len(grid)-2; y++ {
		for x := 1; x <= len(grid[0])-2; x++ {
			if isTreeVisible(&grid, x, y) {
				visibleInnerTrees++
			}
		}
	}

	fmt.Println("Total visible trees:", edgeTreeCount+visibleInnerTrees)
}

func isTreeVisible(grid *[][]int, x int, y int) bool {
	treeHeight := (*grid)[y][x]

	// check visibility from the left
	if treeHeight > slices.Max((*grid)[y][:x]) {
		return true
	}

	// check visibility from the right
	if treeHeight > slices.Max((*grid)[y][x+1:]) {
		return true
	}

	// check visibility from the top
	visible := false
	for topy := 0; topy < y; topy++ {
		if (*grid)[topy][x] >= treeHeight {
			visible = false
			break
		} else {
			visible = true
		}
	}

	if visible {
		return visible
	}

	// check visibility from the bottom
	for bottomy := len(*grid) - 1; bottomy > y; bottomy-- {
		if (*grid)[bottomy][x] >= treeHeight {
			visible = false
			break
		} else {
			visible = true
		}
	}

	return visible
}
