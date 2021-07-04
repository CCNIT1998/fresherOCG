package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

func calculate(slice []string) (result []float64) {
	min, _ := strconv.ParseFloat(slice[0],64)
	max, _ := strconv.ParseFloat(slice[0],64)
	total := 0.0

    for _, v := range slice {
		value, _ := strconv.ParseFloat(v,64)
        if max < value {
            max = value
        }
        if min > value {
            min = value
        }
		total+=value
    }
	average := total/float64(len(slice))
	result = append(result, min, max, average,total)
    return result
}

func main() {  
    data, err := ioutil.ReadFile("text.txt")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
	re := regexp.MustCompile(`\d+\.?\d*`)
	slice := re.FindAllString(string(data), -1)
	
	result := calculate(slice)
	fmt.Println("lenght: ",len(slice))
	fmt.Println("total: ",result[3])
	fmt.Println("Min: ",result[0])
	fmt.Println("Max: ",result[1])
	fmt.Printf("Average: %.2f\n",result[2])
}