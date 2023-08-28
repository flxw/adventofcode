package main

import (
	"adventofcode2022/common"
	"fmt"
)

const PACKET_LENGTH int = 14

func main() {
	var line string = common.ReadWholeFile("input.txt")
	var begin = PACKET_LENGTH - 1
	var end = len(line)

	for ; begin < end; begin++ {
		s := line[begin-(PACKET_LENGTH-1) : begin+1]

		if isSliceUnique(s) {
			break
		}
	}

	fmt.Println("Unique character sequence starts at", begin+1)
}

func isSliceUnique(arr string) bool {
	m := make(map[rune]bool)
	for _, i := range arr {
		_, ok := m[i]
		if ok {
			return false
		}

		m[i] = true
	}

	return true
}
