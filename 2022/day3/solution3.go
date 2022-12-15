package main

import (
	"adventofcode2022/common"
	"fmt"
	"unicode"
)

func main() {
	var priorities []int
	var badgePriorities []int
	var sharedItemsBetweenTriples = make(map[rune]uint8)
	var rucksackIndex = 0

	scoreCalculator := func(line string) {
		middleIndex := int(len(line) / 2)
		leftCompartmentContents := line[:middleIndex]
		rightCompartmentContents := line[middleIndex:]

		sharedItemsBetweenCompartments := make(map[rune]uint8)
		for _, k := range leftCompartmentContents {
			sharedItemsBetweenCompartments[k] |= (1 << 0)
			sharedItemsBetweenTriples[k] |= (1 << rucksackIndex)
		}
		for _, k := range rightCompartmentContents {
			sharedItemsBetweenCompartments[k] |= (1 << 1)
			sharedItemsBetweenTriples[k] |= (1 << rucksackIndex)
		}

		for k, v := range sharedItemsBetweenCompartments {
			a := v&(1<<0) != 0
			b := v&(1<<1) != 0

			if a && b {
				priorities = append(priorities, calculateItemPriority(k))
				break
			}
		}

		if rucksackIndex < 2 {
			rucksackIndex++
		} else {
			for k, v := range sharedItemsBetweenTriples {
				a := v&(1<<0) != 0
				b := v&(1<<1) != 0
				c := v&(1<<2) != 0

				if a && b && c {
					badgePriorities = append(badgePriorities, calculateItemPriority(k))
					break
				}
			}

			rucksackIndex = 0
			sharedItemsBetweenTriples = make(map[rune]uint8)
		}
	}

	common.ReadFileLineByLine("input.txt", scoreCalculator)
	fmt.Println("The combined priority of the mixed-up items is:", common.Sum(priorities...))
	fmt.Println("The combined priority of the badges is:", common.Sum(badgePriorities...))
}

func calculateItemPriority(item rune) int {
	if unicode.IsUpper(item) {
		return int(item) - 38
	} else {
		return int(item) - 96
	}
}
