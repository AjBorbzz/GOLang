package main

import "fmt"

func main() {
	x := 8
	if x > 10 {
		fmt.Println("x is greater than 10")
	} else if x > 5 {
		fmt.Println("x is greater than 5 but less than or equal to 10")
	} else {
		fmt.Println("x is less than or equal to 5")
	}
}
