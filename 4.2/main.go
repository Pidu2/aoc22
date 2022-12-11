package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
		split := strings.Split(scanner.Text(), ",")
		score += checkOverlap(split[0], split[1])
	}
	fmt.Println(score)
}

func checkOverlap(first, second string) int {
	first_split := strings.Split(first, "-")
	second_split := strings.Split(second, "-")
	if convertToInt(first_split[0]) >= convertToInt(second_split[0]) && convertToInt(first_split[1]) <= convertToInt(second_split[1]) || convertToInt(first_split[0]) <= convertToInt(second_split[0]) && convertToInt(first_split[1]) >= convertToInt(second_split[1]) {
		fmt.Printf("%s and %s\n", first, second)
		fmt.Printf("%s >= %s && %s <= %s || %s <= %s && %s >= %s\n", first_split[0], second_split[0], first_split[1], second_split[1], first_split[0], second_split[0], first_split[1], second_split[1])
		return 1
	}
	if convertToInt(first_split[0]) >= convertToInt(second_split[0]) && convertToInt(first_split[0]) <= convertToInt(second_split[1]) || convertToInt(first_split[1]) >= convertToInt(second_split[0]) && convertToInt(first_split[1]) <= convertToInt(second_split[1]) {
		return 1
	}
	return 0
}

func convertToInt(input string) int {
	conv, _ := strconv.Atoi(input)
	return conv
}
