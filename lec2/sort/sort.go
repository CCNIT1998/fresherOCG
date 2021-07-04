package main

import (
	"fmt"
)

func BubbleSort(s []int)([]int){
	swapped := true
	length := len(s)-1
	for swapped {
		swapped = false
		for i := 0; i < length; i++{
			fmt.Println("index: ",i)
			if s[i] > s[i+1]{
				s[i],s[i+1] = s[i+1],s[i]
				swapped = true
			}
		}	
		length --
		fmt.Println("=================================")
	}
	return s 
}


func mergeSort(items []int) []int {
    var num = len(items)
    if num == 1 {
        return items
    }
     
    middle := int(num / 2)
    var (
        left = make([]int, middle)
        right = make([]int, num-middle)
    )
    for i := 0; i < num; i++ {
        if i < middle {
            left[i] = items[i]
        } else {
            right[i-middle] = items[i]
        }
    }
     
    return merge(mergeSort(left), mergeSort(right))
}
 
func merge(left, right []int) (result []int) {
    result = make([]int, len(left) + len(right))
     
    i := 0
    for len(left) > 0 && len(right) > 0 {
        if left[0] < right[0] {
            result[i] = left[0]
            left = left[1:]
        } else {
            result[i] = right[0]
            right = right[1:]
        }
        i++
    }
     
    for j := 0; j < len(left); j++ {
        result[i] = left[j]
        i++
    }
    for j := 0; j < len(right); j++ {
        result[i] = right[j]
        i++
    }
     
    return
}

func main() {
	input := []int{11, 14, 3, 8, 18, 17, 43}
	result := BubbleSort(input)
	result2 := mergeSort(input)
	fmt.Println(result)
	fmt.Println(result2)
}