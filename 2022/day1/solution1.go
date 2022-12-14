package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main1() {
	var FILENAME string = "input1"

	file, err := os.Open(FILENAME)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	currentCapacity := 0
	var elfCapacities []int

	for scanner.Scan() {
		var line = scanner.Text()

		if line == "" {
			elfCapacities = append(elfCapacities, currentCapacity)
			currentCapacity = 0
		} else {
			var parsedLine, _ = strconv.Atoi(line)
			currentCapacity += parsedLine
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elfCapacities)))

	// solution to part 1
	fmt.Println("Maximum elf capacity is: ", elfCapacities[0])

	// solution to part 2
	fmt.Println("Top three elfs carry: ", sum(elfCapacities[0:3]...))
}

func sum(capacities ...int) int {
	result := 0

	for _, number := range capacities {
		result += number
	}

	return result
}
