package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	snacklist := make([]int, 0)
	for scanner.Scan() {
		if scanner.Text() == "" {
			snacklist = append(snacklist, current)
			current = 0
		} else {
			nbr, _ := strconv.Atoi(scanner.Text())
			current += nbr
		}
	}
	sort.Slice(snacklist, func(i, j int) bool {
		return snacklist[i] > snacklist[j]
	})
	result := snacklist[0] + snacklist[1] + snacklist[2]
	fmt.Println(result)
}
