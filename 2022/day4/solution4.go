package main

import (
	"adventofcode2022/common"
	"fmt"
	"strconv"
	"strings"
)

func makeRangeFromStrings(min, max string) []int {
	minInt, _ := strconv.Atoi(min)
	maxInt, _ := strconv.Atoi(max)
	return makeRange(minInt, maxInt)
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func main() {
	var fullyContainedSections int = 0
	var partiallyOverlappingSections int = 0

	lineConsumer := func(line string) {
		pairs := strings.Split(line, ",")
		sectionA := strings.Split(pairs[0], "-")
		sectionB := strings.Split(pairs[1], "-")

		sectionABegin, _ := strconv.Atoi(sectionA[0])
		sectionAEnd, _ := strconv.Atoi(sectionA[1])
		sectionBBegin, _ := strconv.Atoi(sectionB[0])
		sectionBEnd, _ := strconv.Atoi(sectionB[1])

		if (sectionABegin <= sectionBBegin && sectionBEnd <= sectionAEnd) ||
			(sectionBBegin <= sectionABegin && sectionAEnd <= sectionBEnd) {
			fullyContainedSections++
		}

		if (sectionABegin <= sectionBBegin && sectionBBegin <= sectionAEnd) ||
			(sectionBBegin <= sectionABegin && sectionABegin <= sectionBEnd) {
			partiallyOverlappingSections++
		}
	}

	common.ReadFileLineByLine("input1.txt", lineConsumer)
	fmt.Println(fullyContainedSections, "section pairs fully contain each other")
	fmt.Println(partiallyOverlappingSections, "section pairs have partial overlap")
}
