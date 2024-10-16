package main

import "fmt"

func modifyArray(arr [5]int) {
	arr[0] = 100
}

func modifySlice(slice []int) {
	slice[0] = 100
}

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	modifyArray(arr)
	fmt.Println(arr)

	slice := []int{1, 2, 3, 4, 5}
	modifySlice(slice)
	fmt.Println(slice)
}
