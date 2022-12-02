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
	ROCK_Y     = "X"
	PAPER_Y    = "Y"
	SCISSORS_Y = "Z"
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

func calcScore(enemy string, you string) int {
	score := 0
	if you == ROCK_Y {
		score += 1
		switch enemy {
		case ROCK_E:
			score += 3
		case SCISSORS_E:
			score += 6
		}
	} else if you == PAPER_Y {
		score += 2
		switch enemy {
		case ROCK_E:
			score += 6
		case PAPER_E:
			score += 3
		}
	} else if you == SCISSORS_Y {
		score += 3
		switch enemy {
		case PAPER_E:
			score += 6
		case SCISSORS_E:
			score += 3
		}
	}
	return score
}
