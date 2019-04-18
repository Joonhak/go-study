package main

import "fmt"

/*
 * map - `Hash table`을 구현한 자료구조
 */
func main() {
	var _ map[int]string          // 선언
	idMap := make(map[int]string) // 선언
	initMap := map[string]string{ // 선언과 리터럴 초기화
		"key":    "value",
		"like":   "json",
		"python": "dict",
	}

	idMap[0] = "Apple"
	idMap[1] = "Pear"
	idMap[2] = "Peach"

	fmt.Println(idMap)
	delete(idMap, 1) // 삭제 함수
	fmt.Println(idMap)
	fmt.Println(idMap[10])

	fmt.Println(initMap["key"])
	initMap["key"] = "changed value" // `key`가 이미 존재할 때에는 삽입하지 않고 갱신한다.
	fmt.Println(initMap)

	val, exist := initMap["python"]
	if exist {
		fmt.Println(val, exist)
	}

	// key, value 를 이용한 for - range loop 가능
	for k, v := range initMap {
		fmt.Println(k + " : " + v)
	}

}
