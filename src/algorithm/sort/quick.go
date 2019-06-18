package main

import "fmt"

/**
 * Quick sort
 * O(n^2)의 최악, O(n log n)의 평균, 최선 시간복잡도
 * pivot을 선택하고, 선택된 pivot을 기준으로 작은값은 앞으로, 큰 값은 뒤로 정렬한다.
 * 따라서, 어떤 기준으로 pivot을 설정하는지가 성능에 영향을 미친다.
 */
func swap(x *int, y *int) {
	tmp := *x
	*x = *y
	*y = tmp
}

func quickSort(arr []int) []int {
	return nil
}

func main() {
	arr := []int{9, 1, 2, 8, 6, 7, 5, 3, 0, 4}
	fmt.Println(quickSort(arr))
}
