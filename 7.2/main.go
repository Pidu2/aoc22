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

var foldersizes []int

func (n Node) calcSizes() int {
	size := 0
	for _, fsize := range n.files {
		size += fsize
	}
	for _, el := range n.children {
		size += el.calcSizes()
	}
	foldersizes = append(foldersizes, size)
	return size
}

func main() {

	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var rootNode Node
	rootNode.name = "/"
	rootNode.files = make(map[string]int)
	rootNode.children = make(map[string]*Node)

	currentNode := &rootNode

	for scanner.Scan() {
		line := scanner.Text()
		// if root cd -> continue as this is already created
		if line == "$ cd /" {
			continue
		} else if line == "$ cd .." { // if cd upwards -> switch to parent node
			currentNode = currentNode.parent
		} else if strings.HasPrefix(line, "$ cd") { // if cd downwards -> switch to children node
			currentNode = currentNode.children[strings.Split(line, " ")[2]]
		} else if strings.HasPrefix(line, "$ ls") { // if ls -> do nothing
			continue
		} else if strings.HasPrefix(line, "dir") { // inside ls dir list -> create children node
			currentNode.children[strings.Split(line, " ")[1]] = &Node{
				parent:   currentNode,
				name:     strings.Split(line, " ")[1],
				files:    make(map[string]int),
				children: make(map[string]*Node),
			}
		} else { // inside ls file list -> create file entry with size
			split := strings.Split(line, " ")
			name := split[1]
			size, _ := strconv.Atoi(split[0])
			currentNode.files[name] = size
		}
	}

	rootSize := rootNode.calcSizes()
	space_free := 70000000 - rootSize
	space_needed := 30000000 - space_free
	deletesize := rootSize
	for _, size := range foldersizes {
		if size >= space_needed && size < deletesize {
			deletesize = size
		}
	}
	fmt.Println(deletesize)

}
