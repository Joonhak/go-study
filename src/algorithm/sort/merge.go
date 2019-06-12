package main

import "fmt"

/**
 * Merge sort
 * O(n log n)의 시간복잡도
 * 배열을 가장 작은 단위까지 분할 후 병합한다. (재귀)
 * 아래 코드 기준으로 Merge 함수에서 실질적인 정렬이 실행된다.
 */
func main() {
	arr := []int{9, 1, 2, 8, 6, 7, 5, 3, 0, 4}
	fmt.Println(MergeSort(arr))
}

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	middle := len(arr) / 2

	left := MergeSort(arr[:middle])
	right := MergeSort(arr[middle:])

	return Merge(left, right)
}

func Merge(left []int, right []int) []int {
	fmt.Println("left", left, "right", right)
	result := make([]int, 0, len(left) + len(right))

	for len(left) > 0 || len(right) > 0  {
		if len(left) == 0 {
			return append(result, right...)
		}
		if len(right) == 0 {
			return append(result, left...)
		}
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	return result
}