package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("results.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var total_score = 0
	for scanner.Scan() {
		var results = scanner.Bytes()
		var opponent_result = int(results[0]) % 64
		var round_result = (int(results[2]) % 88) * 3
		var my_score int = 0
		//tie
		if round_result == 3 {
			total_score = total_score + opponent_result + round_result
		} else if round_result == 6 {
			//win
			my_score = opponent_result + 1
			if my_score > 3 {
				my_score = my_score % 3
			}
			total_score = total_score + my_score + round_result
		} else {
			//loss
			my_score = opponent_result - 1
			if my_score < 1 {
				my_score = 3
			}
			total_score = total_score + my_score
		}
	}
	fmt.Println(total_score)
}
