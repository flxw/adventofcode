package common

import (
	"bufio"
	"log"
	"os"
)

type consumer func(string)

func ReadFileLineByLine(filename string, fn consumer) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fn(scanner.Text())
	}
}

func Sum(capacities ...int) int {
	result := 0

	for _, number := range capacities {
		result += number
	}

	return result
}
