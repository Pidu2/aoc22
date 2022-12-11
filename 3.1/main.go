package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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
		mid := len(scanner.Text()) / 2
		first, second := scanner.Text()[0:mid], scanner.Text()[mid:]
		common := getCommon(first, second)
		score += getPrio(common)
	}
	fmt.Println(score)
}

func getCommon(first, second string) rune {
	for _, value := range first {
		if strings.Contains(second, string(value)) {
			return value
		}
	}
	return 0
}

func getPrio(input rune) int {
	if input >= 97 {
		return int(input) - 96
	} else {
		return int(input) - 38
	}
}
