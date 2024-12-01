package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Day 6")
	b, errFile := ioutil.ReadFile("input")
	if errFile != nil {
		panic(errFile)
	}
	stream := string(b)

	for i, _ := range stream {
		if i+14 > len(stream) {
			fmt.Println("FAIL")
			break
		}
		buf := stream[i : i+14]

		fakeSet := make(map[string]struct{})

		for _, item := range buf {
			fakeSet[string(item)] = struct{}{}
		}
		if len(fakeSet) == 14 {
			fmt.Println(i + 14)
			break
		}
	}

}
