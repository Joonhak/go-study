package main

import "fmt"

/**
 * Cocktail sort
 * 양방향 Bubble sort라고도 한다.
 * O(n^2)의 최악, 평균 시간 복잡도, O(n)의 최선 시간 복잡도
 * 단방향으로만 진행하는 Bubble sort와 달리
 * 양방향으로 정렬을 진행하기 때문에 Bubble sort보다 빠르다.
 */
func main() {
	arr := []int{9, 1, 2, 8, 6, 7, 5, 3, 0, 4}
	fmt.Println("initial array: ", arr)

	var tmp int

	for i := 0; i < len(arr) / 2; i++ {
		leftIndex := 0
		rightIndex := len(arr) - 1

		for leftIndex <= rightIndex  {
			if arr[leftIndex] > arr[leftIndex + 1] {
				tmp = arr[leftIndex]
				arr[leftIndex] = arr[leftIndex + 1]
				arr[leftIndex + 1] = tmp
			}
			leftIndex++

			if arr[rightIndex] < arr[rightIndex - 1] {
				tmp = arr[rightIndex]
				arr[rightIndex] = arr[rightIndex - 1]
				arr[rightIndex - 1] = tmp
			}
			rightIndex--
		}
	}

	fmt.Println("result: ", arr)
}
