package main

import "fmt"

func sumNumbers(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}
func main() {
	val := sumNumbers(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("The sum : ", val)

}
