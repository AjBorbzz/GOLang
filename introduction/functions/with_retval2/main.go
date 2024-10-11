package main

import "fmt"

func dividedby(a, b int) (int, string) {
	if b == 0 {
		return 0, "Error: division by zero"
	}
	return a / b, ""
}
func main() {
	val, err := dividedby(20, 2)
	if err != "" {
		fmt.Println(err)
	} else {
		fmt.Println("The division : ", val)
	}
}
