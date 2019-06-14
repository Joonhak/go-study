package main

import "fmt"

/**
 * Gnome sort
 * O(n^2)의 최악, O(n)의 최선 시간 복잡도
 * 인접한 원소와 비교하여 정렬하는점은 Bubble sort와 유사하다.
 */
func main() {
	arr := []int{9, 1, 2, 8, 6, 7, 5, 3, 0, 4}
	fmt.Println("initial array: ", arr)

	var tmp int
	i := 1
	for i < len(arr) {
		if arr[i] >= arr[i - 1] {
			i++
		} else {
			tmp = arr[i]
			arr[i] = arr[i - 1]
			arr[i - 1] = tmp

			if i > 1 {
				i--
			}
		}
	}

	fmt.Println("result: ", arr)
}
