package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	ROCK_E     = "A"
	PAPER_E    = "B"
	SCISSORS_E = "C"
	LOSE       = "X"
	DRAW       = "Y"
	WIN        = "Z"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	score := 0
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		score += calcScore(split[0], split[1])
	}
	fmt.Println(score)

}

func calcScore(enemy string, outcome string) int {
	score := 0
	if outcome == DRAW {
		score += 3
		switch enemy {
		case ROCK_E:
			score += 1
		case SCISSORS_E:
			score += 3
		case PAPER_E:
			score += 2
		}
	} else if outcome == WIN {
		score += 6
		switch enemy {
		case ROCK_E:
			score += 2
		case SCISSORS_E:
			score += 1
		case PAPER_E:
			score += 3
		}
	} else if outcome == LOSE {
		switch enemy {
		case ROCK_E:
			score += 3
		case SCISSORS_E:
			score += 2
		case PAPER_E:
			score += 1
		}
	}
	return score
}
