package main

import (
	"adventofcode2022/common"
	"fmt"
	"strconv"
	"strings"
)

func increaseCycle(cycle *int, register *int, sum *int) {
	/*
		uncomment this block for solution to part 1
		if (*cycle-20)%40 == 0 {
			fmt.Println(*cycle, *register, *cycle**register)
			*sum += (*register) * (*cycle)
		}
	*/
	// block for part 2
	spriteStart := *register - 1
	spriteEnd := *register + 1

	if spriteStart <= crtPixelPosition && crtPixelPosition <= spriteEnd {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}

	crtPixelPosition += 1
	crtPixelPosition = crtPixelPosition % 40

	*cycle++

	if (*cycle-1)%40 == 0 {
		fmt.Println()
	}
}

var crtPixelPosition = 0

func main() {
	var cycle = 1
	var register = 1
	var sum = 0

	var lineConsumer = func(line string) {
		lineParts := strings.Split(line, " ")

		switch lineParts[0] {
		case "noop":
			increaseCycle(&cycle, &register, &sum)
		case "addx":
			increaseCycle(&cycle, &register, &sum)
			increaseCycle(&cycle, &register, &sum)
			value, _ := strconv.Atoi(lineParts[1])
			register += value
		}
	}

	common.ReadFileLineByLine("input.txt", lineConsumer)
}
