package main

import "fmt"

func main() {
	numbers := make([]int, 3, 5)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
}
