package main

import "fmt"

/**
 * Selection sort
 * O(n^2) 의 시간 복잡도
 * 배열을 순회하며 가장 작은 원소를 찾아 앞으로 보내는 작업을 반복한다.
 */
func main() {
	arr := [10]int{9, 1, 2, 8, 6, 7, 5, 3, 0, 4}
	fmt.Println("initial array: ", arr)

	var minIndex int
	var tmp int

	for i := 0; i < len(arr); i++ {
		minIndex = i
		for j := i + 1; j < len(arr); j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}

		tmp = arr[i]
		arr[i] = arr[minIndex]
		arr[minIndex] = tmp
	}

	fmt.Println("result: ", arr)
}
