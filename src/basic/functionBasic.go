package main

import "fmt"

/*
 * Go는 자동형변환을 지원하지 않는다.
 * ex) 아래 sum 함수에서 `return type`을 `int`로 변경하면 compile error 발생
 */
func main() {
	// uint = unsigned int ( only positive value )
	var a, b uint = 3, 5

	// 여러값을 return 받을 경우 차례로 변수에 담을 수 있다.
	result, origin1, origin2 := addition(a, b)

	// 변수에 할당 가능
	multiplication := func(i1 int, i2 int) int {
		return i1 * i2
	}

	fmt.Printf("Result: %d, Origin1: %d, Origin2: %d \n", result, origin1, origin2)
	fmt.Printf("              Subtraction Reulst: %d \n", calculate(func(x int, y int) int { return x - y }, 5, 3))
	fmt.Printf("           Multiplication Result: %d \n", multiplication(10, 10))
	fmt.Printf("                 Division Reulst: %d \n", division(func(x int, y int) int { return x / y }, 10, 5))
}

// 한번에 여러 값 return 가능
func addition(i1 uint, i2 uint) (uint, uint, uint) {
	return i1 + i2, i1, i2
}

// 함수를 `parameter`로 전달받고, 전달할 수 있음
func calculate(f func(int, int) int, i1 int, i2 int) int {
	result := f(i1, i2)
	return result
}

// `type`문을 통해 함수의 원형 지정 가능
type calc func(int, int) int

// `type`문을 통해 선언한 함수의 원형 사용
func division(f calc, i1 int, i2 int) int {
	result := f(i1, i2)
	return result
}
