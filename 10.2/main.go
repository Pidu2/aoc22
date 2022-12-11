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
	for scanner.Scan() {
		line := scanner.Text()
		spl := strings.Split(line, " ")
		op := spl[0]
		draw(cycle, x)
		if op == "addx" {
			cycle += 1
			draw(cycle, x)
			cycle += 1
			nbr, _ := strconv.Atoi(spl[1])
			x += nbr
		} else {
			cycle += 1
		}
	}
	fmt.Println()
}

func draw(cycle int, x int) {
	cycle_pos := (cycle - 1) % 40
	if cycle_pos == 0 {
		fmt.Println()
	}
	if cycle_pos >= x-1 && cycle_pos <= x+1 {
		fmt.Printf("#")
	} else {
		fmt.Printf(".")
	}
}
