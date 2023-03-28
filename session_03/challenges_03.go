package main

import (
	"fmt"
)

func dup_count(list []string) map[string]int {

	duplicate_frequency := make(map[string]int)

	for _, item := range list {
		// check if the item/element exist in the duplicate_frequency map
		_, exist := duplicate_frequency[item]

		if exist {
			duplicate_frequency[item] += 1 // increase counter by 1 if already in the map
		} else {
			duplicate_frequency[item] = 1 // else start counting from 1
		}
	}
	return duplicate_frequency
}

func main() {

	var inputSlice = []string{}
	var input = "selamat malam"

	fmt.Println("Mohammad Fatoni Anggris (1955617840-17)")
	fmt.Println("-----------------------------------------")

	chars := []rune(input)
	for i := 0; i < len(chars); i++ {
		char := string(chars[i])
		inputSlice = append(inputSlice, char)
		fmt.Println(char)
	}

	removeDuplicateValuesSlice := dup_count(inputSlice)
	fmt.Println(removeDuplicateValuesSlice)

}
