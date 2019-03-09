package main

import "fmt"

/*
 * 배열 선언 방법
 *      // int 3개를 갖는 배열
 *   1. var arr [3]int
 *      // 배열 선언과 동시에 초기화
 *   2. var arr = [3]int{1, 2, 3}
 *      // 선언과 동시에 초기화 시 배열 크기 생략 가능
 * 2-2. var arr = []int{1, 2, 3}
 *      // 배열 선언과 동시에 초기화 시 var 키워드 생략 가능
 *   3. arr := []int{1, 2, 3}
 *      // 다차원 배열 선언, 초기화 시 아래와 같이 작성
 *   4. var arr = [3][3]int{ {1, 2, 3}, {1, 2, 3}, }
 *
 */
func main() {
	a := []int{1, 2, 3}
	b := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}, // 마지막에도 `comma`를 넣어줘야한다.
	}

	for _, i := range a {
		println(i)
	}
	for i := range b {
		for _, j := range b[i] {
			println(j)
		}
	}

	fmt.Printf("a array: %v\n", a)
	fmt.Printf("b array: %v\n", b)
}
