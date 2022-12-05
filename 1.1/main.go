package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	current := 0
	biggest := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			if current > biggest {
				biggest = current
			}
			current = 0
		} else {
			nbr, _ := strconv.Atoi(scanner.Text())
			current += nbr
		}
	}
	fmt.Println(biggest)
}
