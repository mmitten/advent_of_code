package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func doesOverlap(pair1, pair2 []int) bool {
	// sort
	if pair1[0] > pair2[0] {
		pair1, pair2 = pair2, pair1
	}
	return pair1[1] >= pair2[0]
}

func get_ranges(input string) [][]int {
	var split = strings.Split(input, ",")
	var output [][]int
	for _, text_range := range split {
		var range_split = strings.Split(text_range, "-")
		var int_range []int
		for _, text_int := range range_split {
			parsed, parse_err := strconv.Atoi(text_int)
			if parse_err != nil {
				log.Fatal(parse_err)
			}
			int_range = append(int_range, parsed)
		}
		output = append(output, int_range)
	}
	return output
}

func main() {
	file, err := os.Open("ranges.txt")
	if err != nil {
		log.Fatal(err)
	}
	var part_1_count = 0
	var part_2_count = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var ranges = get_ranges(scanner.Text())
		if ranges[1][0] >= ranges[0][0] && ranges[1][1] <= ranges[0][1] {
			part_1_count++
		} else if ranges[0][0] >= ranges[1][0] && ranges[0][1] <= ranges[1][1] {
			part_1_count++
		}

		if doesOverlap(ranges[0], ranges[1]) {
			part_2_count++
		}
	}
	fmt.Println("Part 1:", part_1_count)
	fmt.Println("Part 2:", part_2_count)
}
