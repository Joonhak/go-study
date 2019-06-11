package main

import "fmt"

func main() {
	arr := [10]int{9, 1, 2, 8, 6, 7, 5, 3, 0, 4}
	fmt.Println("initial array: ", arr)

	var min int
	var tmp int

	for i := 0; i < len(arr); i++ {
		min = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		tmp = arr[i]
		arr[i] = arr[min]
		arr[min] = tmp
	}

	fmt.Println("result: ", arr)
}
