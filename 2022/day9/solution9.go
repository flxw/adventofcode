package main

import (
	"adventofcode2022/common"
	"fmt"
	"strconv"
	"strings"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Coordinate struct {
	x int
	y int
}

func (c *Coordinate) DistanceGreaterThan(to *Coordinate, than int) bool {
	xDist := Abs(to.x - c.x)
	yDist := Abs(to.y - c.y)

	return xDist > than || yDist > than
}

func main() {
	var head = Coordinate{0, 0}
	var tail = Coordinate{0, 0}
	tailCoordinateVisits := map[Coordinate]int{}

	tailCoordinateVisits[Coordinate{0, 0}] = 1

	var lineConsumer = func(line string) {
		lineParts := strings.Split(line, " ")
		direction := lineParts[0]
		stepCount, _ := strconv.Atoi(lineParts[1])

		for i := 0; i < stepCount; i++ {
			switch direction {
			case "R":
				head.x += 1
			case "L":
				head.x -= 1
			case "U":
				head.y += 1
			case "D":
				head.y -= 1
			}

			if tail.DistanceGreaterThan(&head, 1) {
				switch direction {
				case "R":
					tail.x += 1
					tail.y = head.y
				case "L":
					tail.x -= 1
					tail.y = head.y
				case "U":
					tail.y += 1
					tail.x = head.x
				case "D":
					tail.y -= 1
					tail.x = head.x
				}
				tailCoordinateVisits[tail] = tailCoordinateVisits[tail] + 1
			}
		}
	}

	common.ReadFileLineByLine("input.txt", lineConsumer)
	fmt.Println(tailCoordinateVisits)
	fmt.Println(len(tailCoordinateVisits))
}
