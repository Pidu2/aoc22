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
	i := 0
	var group []string
	for scanner.Scan() {
		if i%3 == 0 && i != 0 {
			common := getCommon(group)
			score += getPrio(common)
			group = nil
		}
		group = append(group, scanner.Text())
		i++
	}
	common := getCommon(group)
	score += getPrio(common)
	fmt.Println(score)
}

func getCommon(group []string) rune {
	fmt.Println(group)
	for _, value := range group[0] {
		if strings.Contains(group[1], string(value)) && strings.Contains(group[2], string(value)) {
			fmt.Println(string(value))
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
