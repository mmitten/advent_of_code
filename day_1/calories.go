package main

import (
	"os"
	"log"
	"bufio"
	"strconv"
	"fmt"
)
	
func main() {
	file, err := os.Open("data.txt")
	if err != nil{
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var calorie_count = 0
	var highest_calories = 0
	for scanner.Scan() {
		item := scanner.Text()
		if item == ""{
			if calorie_count > highest_calories {
				highest_calories = calorie_count
			}
			calorie_count = 0
		} else {
			single, err := strconv.Atoi(item)
			if err != nil{
				log.Fatal(err)
			}
			calorie_count += single
		}
	}
	fmt.Println(highest_calories)
}

