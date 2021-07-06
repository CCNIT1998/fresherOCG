package main

import (
	"fmt"
)


func QuickSort(s []int, left, right int) []int {
	
	if left < right {
		var p int
		s, p = partition(s, left, right)
		s = QuickSort(s, left, p-1)
		s = QuickSort(s, p+1, right)
	}
	return s
}

func partition(s []int, left, right int)([]int, int){
	pivot := s[right]
	i := left
	for j := left; j < right; j++ {
		if s[j] < pivot {
			s[i], s[j] = s[j], s[i]
			i++
		}
	}
	s[i], s[right] = s[right], s[i]
	return s, i
}


func main() {
	input := []int{4,33, 1, 3,11, 8, 9,6,22, 7, 5}
	a := QuickSort(input,0,len(input)-1)
	fmt.Println(a)
}