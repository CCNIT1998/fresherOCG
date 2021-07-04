package main

import (
	"fmt"
	"math"
)

func checkPrime(n int){
	for i := 2; i <= math.Sqrt(n); i++{
        if(n % i == 0){
            // count++;
        }
    }
}

func main() {
	fmt.Println()
}