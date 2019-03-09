package main

/*
 * Closure: 함수 밖에 있는 변수를 참조하는 함수값
 */
func main() {
	/*
	 * 이 때 i는 nextValue()가 반환하는 함수의 밖에 있는 값
	 * nextValue() = func() int {
	 *     i++
	 *     return i
	 * }
	 */
	next := nextValue()
	println(next()) // -> 1
	println(next()) // -> 2
	println(next()) // -> 3

	anotherNext := nextValue() // nextValue() 안의 i가 0으로 새로 초기화됨
	println(anotherNext()) // -> 1
	println(anotherNext()) // -> 2

	println(next()) // -> 4 ( 유지됨 )
}

func nextValue() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}