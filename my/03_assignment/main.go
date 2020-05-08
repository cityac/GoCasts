package main

import "fmt"

func main() {
	printOddsEvens(10)
}

func printOddsEvens(size int) {
	numbers := []int{}

	for i := 0; i < size; i++ {
		numbers = append(numbers, i)
	}

	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Println(fmt.Sprintf("%d is even", number))
		} else {
			fmt.Println(fmt.Sprintf("%d is odd", number))
		}
	}
}
