package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
)

type Tree struct {
	size    int
	visible bool
	score   int
}

type TreeMap map[image.Point]Tree

func LoadGrid() (TreeMap, error) {
	file, errFile := os.Open("tree_grid")
	if errFile != nil {
		return nil, errFile
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var grid = make(TreeMap)
	var dy int = 0
	var maxX int = 0
	for scanner.Scan() {
		holder := scanner.Bytes()
		for dx, item := range holder {
			height := int(item) - 48
			tree := Tree{height, false, 1}
			grid[image.Pt(dx, dy)] = tree
			maxX = dx
		}
		dy += 1
	}
	var maxY = dy - 1
	fmt.Println(maxX, maxY)
	for point, tree := range grid {

		for _, d := range []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
			for i := 1; ; i++ {
				if nextTree, ok := grid[point.Add(d.Mul(i))]; !ok {
					tree.visible = true
					tree.score *= (i - 1)
					grid[point] = tree
					break
				} else if nextTree.size >= tree.size {
					tree.score *= i
					grid[point] = tree
					break
				}
			}
		}
	}
	return grid, nil
}

func main() {
	data, errData := LoadGrid()
	if errData != nil {
		panic(errData)
	}

	var part1 int = 0
	var part2 int = 0
	for _, tree := range data {
		if tree.visible {
			part1++
		}
		if tree.score > part2 {
			part2 = tree.score
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
