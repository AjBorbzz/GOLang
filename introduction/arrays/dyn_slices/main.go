package main 

import "fmt"

func main() {
	numbers:=[]int{1,2,3}
	// append in numbers
	numbers=append(numbers,4,5,6,7)
	fmt.Println(numbers)
}