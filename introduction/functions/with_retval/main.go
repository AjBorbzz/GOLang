package main

import "fmt"

func sum(a, b int) int {
	return a + b
}
func main() {
	var val = sum(5, 9)
	fmt.Println("The sum : ", val)
}
