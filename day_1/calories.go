package main

import (
	"os"
	"log"
	"bufio"
	"strconv"
	"fmt"
	"sort"
)
	
func main() {
	file, err := os.Open("data.txt")
	if err != nil{
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var calorie_count = 0
	var highest_calories []int;
	for scanner.Scan() {
		item := scanner.Text()
		if item == ""{
			highest_calories = append(highest_calories, calorie_count)
			calorie_count = 0
		} else {
			single, err := strconv.Atoi(item)
			if err != nil{
				log.Fatal(err)
			}
			calorie_count += single
		}
	}
	
	sort.Sort(sort.Reverse(sort.IntSlice(highest_calories)))
	fmt.Println(highest_calories[0] + highest_calories[1] + highest_calories[2])
}

