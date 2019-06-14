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
	leftIndex := 0
	rightIndex := len(arr) - 1

	for i := 0; i < len(arr) / 2; i++ {
		for i := leftIndex; i < rightIndex; i++ {
			if arr[i] > arr[i + 1] {
				tmp = arr[i]
				arr[i] = arr[i + 1]
				arr[i + 1] = tmp
			}
		}

		for j := rightIndex; j > leftIndex; j-- {
			if arr[j] < arr[j - 1] {
				tmp = arr[j]
				arr[j] = arr[j - 1]
				arr[j - 1] = tmp
			}
		}

		leftIndex++
		rightIndex--
	}

	fmt.Println("result: ", arr)
}
