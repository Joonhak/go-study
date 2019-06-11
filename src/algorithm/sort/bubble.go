package main

import "fmt"

/**
 * Bubble sort
 * O(n^2) 의 시간복잡도
 * 배열의 i, i + 1 번째(다음 인덱스) 원소를 비교하여
 * 큰 원소를 뒤로, 작은 원소를 앞으로 보내는 작업을 반복한다.
 * 느리지만 간단한 코드로 작성이 가능하다.
 */
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


