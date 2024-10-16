package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, value := range numbers {
		fmt.Println("Index", i, "Value:", value)
	}
	fmt.Println(numbers[3:5])
}
