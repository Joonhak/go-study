package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i []int

	if i == nil {
		// 길이와 용량을 할당하지 않으면 `nil slice`를 생성하며, `nil`과 같다.
		fmt.Println("a is nil") // print
	}

	i = []int{1, 2, 3, 4, 5, 6, 7}
	i[1] = 10

	// type, length ( 생략 불가 ), capacity ( 최대 길이? 생략가능 )
	s := make([]string, 2, 5)
	s2 := []string{"a", "b", "c", "d", "e"}

	for i := 0; i < len(s); i++ {
		s[i] = "aaa" + strconv.Itoa(i)
	}

	// 최대 길이가 초과되게 append 시, 최대 용량을 기존 대비 2배로 설정한다.
	// like java - ArrayList<T>
	s = append(s, "asd1", "asd2", "asd3", "asd4", "asd5")
	s = append(s, s2...) // 뒤에 붙는 ...는 spread operator

	fmt.Println(i[2:]) // 2 (index) ~ 마지막 까지 출력
	fmt.Println(s)

	source := []int{0, 1, 2}
	target := make([]int, len(source), cap(source)*2)

	copy(target, source)

	fmt.Println(target)
	fmt.Println(source)

	fmt.Println(len(source), cap(source), len(target), cap(target))
}
