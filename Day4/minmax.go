package main

import (
	"fmt"
)

func findMinMax(numbers []int) (int, int) {
	min := numbers[0]
	max := numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return min, max
}

func main() {
	var numbers []int
	for i := 0; i < 6; i++ {
		var num int
		fmt.Printf("Input number %d: ", i+1)
		fmt.Scan(&num)
		numbers = append(numbers, num)
	}

	min, max := findMinMax(numbers)
	fmt.Printf("Minimum value: %d\n", min)
	fmt.Printf("Maximum value: %d\n", max)
}
