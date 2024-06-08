package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Leaf struct {
	Children []*Leaf
	Type     string
	Size     int
	Name     string
	Parent   *Leaf
}

func (l *Leaf) GetChildByName(name string) *Leaf {
	for _, child := range l.Children {
		if child.Name == name {
			return child
		}
	}
	return &Leaf{}
}

func (l *Leaf) GetParent() *Leaf {
	return l.Parent
}

func (l *Leaf) GetDirSize() int {
	if l.Type == "file" {
		return 0
	} else {
		sum := 0
		for _, child := range l.Children {
			if child.Type == "file" {
				sum = sum + child.Size
			}
		}
		return sum
	}
}

func (l *Leaf) ComputeSize() {
	if len(l.Children) == 0 {
		return
	}
	for _, child := range l.Children {
		if child.Type == "dir" && child.Size == 0 {
			child.ComputeSize()
		}
	}
	for _, child := range l.Children {
		l.Size = l.Size + child.Size
	}
	return
}

func (l *Leaf) GenerateSizeMap(sizeSlice *[]int) {
	if l.Type == "file" {
		return
	} else {
		*sizeSlice = append(*sizeSlice, l.Size)
		for _, child := range l.Children {
			child.GenerateSizeMap(sizeSlice)
		}
	}
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
	last := &root
	for i, line := range data {
		if i == 0 {
			continue
		}
		cmds := strings.Split(line, " ")
		if cmds[0] == "$" {
			if cmds[1] == "cd" {
				if cmds[2] != ".." {
					last = last.GetChildByName(cmds[2])
				} else {
					last = last.GetParent()
				}
			}
		} else if cmds[0] == "dir" {
			newDir := Leaf{Name: cmds[1], Type: cmds[0], Parent: last}
			last.Children = append(last.Children, &newDir)
		} else {
			sizeInt, _ := strconv.Atoi(cmds[0])
			newFile := Leaf{Name: cmds[1], Type: "file", Size: sizeInt, Parent: last}
			last.Children = append(last.Children, &newFile)
		}
	}
	root.ComputeSize()
	return root
}

func main() {
	fmt.Println("Day 7")
	data := readFile()
	root := buildTree(data)
	var part1 []int
	root.GenerateSizeMap(&part1)
	var total int
	for _, size := range part1 {
		if size < 100000 {
			total = total + size
		}
	}

	fmt.Println(total)
}
