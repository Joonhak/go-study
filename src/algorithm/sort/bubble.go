package main

import "fmt"

func main() {
	arr := [10]int{9, 1, 2, 8, 6, 7, 5, 3, 0, 4}
	fmt.Println("initial array: ", arr)

	//tmp := 0
	var tmp int

	for i := 0; i < len(arr) - 1; i++ {
		for j := 0; j < len(arr) - 1 - i; j++ {
			if arr[j] < arr[j + 1] {
				tmp = arr[j + 1]
				arr[j + 1] = arr[j]
				arr[j] = tmp
			}
		}
	}

	fmt.Println("result: ", arr)
}


