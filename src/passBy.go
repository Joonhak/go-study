package main

/*
 * Java, Javascript 등의 언어만 사용하며 깊게 생각해본적 없었던 부분..
 * 그냥 전달하면 값을 `복사` 하고,
 * &연산자를 통해 주소값을 전달하면 원본 객체(?)를 수정할 수 있다.
 */
func main() {
	str := "Test string" // var를 생략하고 초기화할 수 있는 := 연산자

	passByValue(str)
	println("pass value :" + str)
	// -> Test string

	passByReference(&str)
	println("pass reference :" + str)
	// -> Changed value!
}

func passByValue(str string) {
	str = "Changed value!"
}

func passByReference(str *string) {
	*str = "Changed value!"
}

