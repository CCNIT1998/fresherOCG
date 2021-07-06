package main

import (
	"fmt"
)

func BubbleSort(s []int) []int {
	swapped := true
	length := len(s) - 1
	for swapped {
		swapped = false
		for i := 0; i < length; i++ {
			fmt.Println("index: ", i)
			if s[i] > s[i+1] {
				s[i], s[i+1] = s[i+1], s[i]
				swapped = true
			}
		}
		length--
		fmt.Println("=================================")
	}
	return s
}

func main() {
	input := []int{4,33, 1, 3,11, 8, 9,6,22, 7, 5}
	result := BubbleSort(input)
	fmt.Println(result)

}
