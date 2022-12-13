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
	stackz := make(map[int][]string)
	ops := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || line[1] == '1' {
			ops = true
			continue
		}
		if !ops {
			for i := 1; i < len(line); i += 4 {
				if line[i] == ' ' {
					continue
				}
				if _, ok := stackz[i/4]; !ok {
					stackz[i/4] = []string{string(line[i])}
				} else {
					stackz[i/4] = append(stackz[i/4], string(line[i]))
				}
			}
		} else {
			nbr, from, to := getOperators(line)
			move(stackz, nbr, from, to)
		}

		// if contains [] -> stack part
		// if does not start with move -> skip
		// else do stuff
	}
	for i := 0; i < len(stackz); i++ {
		fmt.Printf("%s", stackz[i][0])
	}
	fmt.Println()
}

func getOperators(input string) (int, int, int) {
	s := strings.Split(input, " ")
	nbr, _ := strconv.Atoi(s[1])
	from, _ := strconv.Atoi(s[3])
	to, _ := strconv.Atoi(s[5])
	return nbr, from, to
}

func move(stackz map[int][]string, nbr int, from int, to int) {
	var tomove []string
	for i := 0; i < nbr; i++ {
		tomove = append(tomove, stackz[from-1][i])
	}
	stackz[from-1] = stackz[from-1][nbr:]
	stackz[to-1] = append(tomove, stackz[to-1]...)
}
