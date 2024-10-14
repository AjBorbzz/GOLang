package main

import "fmt"

func main() {
	var numbers [5]int
	for i := 0; i < len(numbers); i++ {
		numbers[i] = i * 2
	}
	fmt.Println(numbers)
}
