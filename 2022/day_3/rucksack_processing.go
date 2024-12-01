package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func get_score(character rune) int {
	var int_char = int(character)
	if int(int_char) > 96 {
		return int_char - 96
	} else {
		return (int_char - 38)
	}
}

func score_rucksack(data []byte) int {
	rucksack := string(data)
	sideOne := rucksack[len(rucksack)/2:]
	sideTwo := rucksack[:len(rucksack)/2]
	for _, c := range sideOne {
		if strings.Contains(sideTwo, string(c)) {
			return get_score(c)
		}
	}
	return 0
}

func find_badge(data [][]byte) int {
	var rucksacks []string
	for _, byte_arr := range data {
		rucksacks = append(rucksacks, string(byte_arr))
	}
	for _, c := range rucksacks[0] {
		if strings.Contains(rucksacks[1], string(c)) && strings.Contains(rucksacks[2], string(c)) {
			return get_score(c)
		}
	}
	return 0
}

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	var priority_score int = 0
	var badge_priority int = 0
	var rucksack_group [][]byte
	for scanner.Scan() {
		priority_score += score_rucksack(scanner.Bytes())
		rucksack_group = append(rucksack_group, scanner.Bytes())
		if len(rucksack_group) == 3 {
			badge_priority += find_badge(rucksack_group)
			rucksack_group = nil
		}
	}
	fmt.Println("Part 1:", priority_score)
	fmt.Println("Part 2:", badge_priority)
}
