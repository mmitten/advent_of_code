package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Instruction struct {
	MoveCount   int
	SourceStack int
	DestStack   int
}

func processLine(data string) Instruction {
	s := strings.Split(data, " ")
	if len(s) < 6 {
		return Instruction{}
	}
	moveCount, errMoveCount := strconv.Atoi(s[1])
	if errMoveCount != nil {
		panic(errMoveCount)
	}

	souceStack, errSourceStack := strconv.Atoi(s[3])
	if errSourceStack != nil {
		panic(errSourceStack)
	}

	destStack, errDestStack := strconv.Atoi(s[5])
	if errDestStack != nil {
		panic(errDestStack)
	}

	steps := Instruction{
		MoveCount:   moveCount,
		SourceStack: souceStack - 1,
		DestStack:   destStack - 1,
	}
	return steps
}

func get_format_and_instructions() ([][]string, []Instruction) {
	file, errFile := os.Open("data")
	if errFile != nil {
		fmt.Println(errFile)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var theStacks = make([][]string, 9)
	var instructionSet []Instruction
	var lineNumber = 0
	for fileScanner.Scan() {
		lineNumber++
		data := fileScanner.Text()
		if lineNumber < 10 {
			for i, char := range data {
				if char == ' ' || char == ']' || char == '[' {
					continue
				}
				listIdent := (i + 3) / 4
				if char > 64 {
					theStacks[listIdent-1] = append(theStacks[listIdent-1], string(char))
				}
			}
		} else {
			instruction := processLine(data)
			instructionSet = append(instructionSet, instruction)
		}
	}
	for _, stack := range theStacks {
		slices.Reverse(stack)
	}
	return theStacks, instructionSet
}

func process_part_1(data [][]string, instructions []Instruction) {
	for _, inst := range instructions {
		for i := 0; i < inst.MoveCount; i++ {
			val := data[inst.SourceStack][len(data[inst.SourceStack])-1]
			data[inst.SourceStack] = data[inst.SourceStack][:len(data[inst.SourceStack])-1]
			data[inst.DestStack] = append(data[inst.DestStack], val)
		}
	}
	fmt.Print("Part 1: ")
	for _, stack := range data {
		fmt.Print(stack[len(stack)-1])
	}
	fmt.Println()
}

func process_part_2(data [][]string, instructions []Instruction) {
	for _, inst := range instructions {
		bottomOfStack := len(data[inst.SourceStack]) - (inst.MoveCount)
		if bottomOfStack < 0 {
			bottomOfStack = 0
		}
		topOfStack := len(data[inst.SourceStack])
		if topOfStack < 0 {
			topOfStack = 0
		}
		vals := data[inst.SourceStack][bottomOfStack:topOfStack]
		fmt.Println(inst.MoveCount, vals, data[inst.SourceStack], data[inst.DestStack])
		data[inst.SourceStack] = data[inst.SourceStack][:bottomOfStack]
		data[inst.DestStack] = append(data[inst.DestStack], vals...)
		fmt.Println(vals, data[inst.SourceStack], data[inst.DestStack])
	}
	fmt.Print("Part 2: ")
	for _, stack := range data {
		topOfFinalStack := len(stack) - 1
		if topOfFinalStack < 0 {
			fmt.Print(" ")
		} else {
			fmt.Print(stack[topOfFinalStack])
		}
	}
}

func main() {
	fmt.Println("Day 5")
	data1, instructions1 := get_format_and_instructions()
	process_part_1(data1, instructions1)
	data2, instructions2 := get_format_and_instructions()
	process_part_2(data2, instructions2)
}
