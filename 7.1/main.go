package main

import (
	"bufio"
	"log"
	"os"
)

type Node struct {
	parent   *Node
	children []*Node
	name     string
	files    map[string]int
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
	}
}
