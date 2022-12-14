package main

import (
	"adventofcode2022/common"
	"fmt"
)

func main() {
	// Opponent: A for Rock, B for Paper, and C for Scissors
	// Us: X for Rock, Y for Paper, and Z for Scissors

	score := 0
	scoreCalculator := func(line string) {
		opponentMove := rune(line[0])
		ourMove := rune(line[2]) - 23 // align both to same lettering

		if (opponentMove == 'A' && ourMove == 'B') ||
			(opponentMove == 'B' && ourMove == 'C') ||
			(opponentMove == 'C' && ourMove == 'A') {
			score += 6
		} else if opponentMove == ourMove {
			score += 3
		}

		switch ourMove {
		case 'A':
			score += 1
		case 'B':
			score += 2
		case 'C':
			score += 3
		}
	}

	common.ReadFileLineByLine("input1.txt", scoreCalculator)
	fmt.Println("Our calculated score for the first part is:", score)

	// Part 2 - second column has a different meaning
	// X means you need to lose
	// Y means you need to end the round in a draw, and
	// Z means you need to win. Good luck!"
	score = 0
	scoreCalculator = func(line string) {
		opponentMove := rune(line[0])
		supposedOutcome := rune(line[2])

		switch supposedOutcome {
		case 'X':
			switch opponentMove {
			case 'A':
				score += 3
			case 'B':
				score += 1
			case 'C':
				score += 2
			}
		case 'Y':
			score += 3

			switch opponentMove {
			case 'A':
				score += 1
			case 'B':
				score += 2
			case 'C':
				score += 3
			}
		case 'Z':
			score += 6

			switch opponentMove {
			case 'A':
				score += 2
			case 'B':
				score += 3
			case 'C':
				score += 1
			}
		}
	}

	common.ReadFileLineByLine("input1.txt", scoreCalculator)
	fmt.Println("Our calculated score for the second part is:", score)
}
