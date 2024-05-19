package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Leaf struct {
	Children []Leaf
	Type     string
	Size     int
	Name     string
	Parent   *Leaf
}

func (l *Leaf) GetChildByName(name string) Leaf {
	for _, child := range l.Children {
		if child.Name == name {
			return child
		}
	}
	return Leaf{}
}

func (l *Leaf) GetParent() Leaf {
	return *l.Parent
}

func readFile() []string {
	file, errFile := os.Open("input")
	if errFile != nil {
		panic(errFile)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var instructions []string
	for fileScanner.Scan() {
		instructions = append(instructions, fileScanner.Text())
	}
	return instructions
}

func buildTree(data []string) Leaf {
	root := Leaf{Name: "/", Type: "dir"}
	last := root
	for i, line := range data {
		fmt.Println(last, root)
		if i == 0 {
			continue
		}
		cmds := strings.Split(line, " ")
		if cmds[0] == "$" {
			if cmds[1] == "cd" {
				if cmds[2] != ".." {
					last = last.GetChildByName(cmds[1])
				} else {
					last = last.GetParent()
				}
			}
		} else if cmds[0] == "dir" {
			last.Children = append(last.Children, Leaf{Name: cmds[1], Type: cmds[0], Parent: &last})
		} else {
			sizeInt, _ := strconv.Atoi(cmds[0])
			last.Children = append(last.Children, Leaf{Name: cmds[1], Type: "file", Size: sizeInt, Parent: &last})
		}
	}
	return root
}

func main() {
	fmt.Println("Day 7")
	data := readFile()
	root := buildTree(data)
	fmt.Println(root)
}
