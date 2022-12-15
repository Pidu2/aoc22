package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	parent   *Node
	children map[string]*Node
	name     string
	files    map[string]int
}

func (n Node) show() {
	fmt.Println("Name ", n.name)
	fmt.Printf("Children: %v\n", n.children)
	fmt.Printf("Parent: %v\n", n.parent)
	fmt.Printf("files: %v\n", n.files)
	fmt.Println("#############")
}

func main() {
	/*
			for line in lines:
			  if line startswith $ cd <$name>
			    currentnode = children[<$name>]
			  if line startswith $ cd <..>:
			    currentnode = currentnode.parent
		      if line startswith $ ls: skip
			  aues andere in ls:
			  	if startswith number:
				  currentnode.files[<filename>] = size
				if startswith dir:
				  currentnode.children.append(node(parent=current, children=nil, name="$name", files=nil)
			recursively loop through nodes and return their size
	*/

	file, err := os.Open("input_test")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var currentNode Node

	for scanner.Scan() {
		line := scanner.Text()
		currentNode.show()
		fmt.Println(line)
		if line == "$ cd /" {
			currentNode.name = "/"
			currentNode.files = make(map[string]int)
			currentNode.children = make(map[string]*Node)
		} else if line == "$ cd .." {
			currentNode = *currentNode.parent
		} else if strings.HasPrefix(line, "$ cd") {
			currentNode = *currentNode.children[strings.Split(line, " ")[2]]
		} else if strings.HasPrefix(line, "$ ls") {
			continue
		} else if strings.HasPrefix(line, "dir") {
			currentNode.children[strings.Split(line, " ")[1]] = &Node{
				parent:   &currentNode,
				name:     strings.Split(line, " ")[1],
				files:    make(map[string]int),
				children: make(map[string]*Node),
			}
		} else {
			split := strings.Split(line, " ")
			name := split[1]
			size, _ := strconv.Atoi(split[0])
			currentNode.files[name] = size
		}
	}
}
