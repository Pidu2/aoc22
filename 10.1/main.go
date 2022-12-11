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
	cycle := 1
	x := 1
	signalstrength := 0
	for scanner.Scan() {
		line := scanner.Text()
		spl := strings.Split(line, " ")
		op := spl[0]
		signalstrength += test(cycle, x)
		if op == "addx" {
			cycle += 1
			signalstrength += test(cycle, x)
			cycle += 1
			nbr, _ := strconv.Atoi(spl[1])
			x += nbr
		} else {
			cycle += 1
		}
	}
	fmt.Println("\nSUM: ", signalstrength)
}

func test(cycle int, x int) int {
	if (cycle+20)%40 == 0 {
		fmt.Printf("Signal strength: %d (cycle %d * x %d)\n", (cycle * x), cycle, x)
		return cycle * x
	}
	return 0
}
