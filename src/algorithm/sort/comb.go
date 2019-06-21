package main

import "fmt"

/**
 * Comb sort
 * O(n^2)의 최악, O(n log n)의 최선 시간복잡도
 * Bubble sort의 리스트 끝의 작은값을 (거북이라고도 부른다)
 * 제거하는 방법으로 성능을 개선한 정렬방법이다.
 * Bubble sort가 i와 i + 1을 비교한다면
 * Comb sort는 i와 i + gap (gap은 변하는 값이고 작아진다) 을 비교한다.
 */
func main() {
	arr := [10]int{9, 1, 2, 8, 6, 7, 5, 3, 0, 4}
	fmt.Println("initial array: ", arr)

	tmp := 0
	gap := len(arr)

	for gap > 1 {
		gap = gap * 10 / 13

		for i := 0; i < len(arr) - gap; i++ {
			if arr[i] > arr[i + gap] {
				tmp = arr[i]
				arr[i] = arr[i + gap]
				arr[i + gap] = tmp
			}
		}

	}

	fmt.Println("result: ", arr)
}
