package main

import (
	"fmt"
	"os"
)

func main() {
	f := openFile("D:\\Invalid_file.txt") // "D:\\backup\\기본적인 정보.txt"
	bytes := make([]byte, 512)
	if f != nil {
		f.Read(bytes)
		fmt.Printf("%s\n", bytes)
	}
}

// panic -> 현재 함수를 즉시 멈추고 함수내의 defer 함수들을 실행시키고 함수를 종료시킴. 상위 함수에도 똑같이 적용됨.
func openFile(filename string) *os.File {
	defer func() { // recover 함수, `panic`이 실행되어도 비정상 종료되지 않는다.
		if r := recover(); r != nil {
			fmt.Println("can not open File", r)
		}
	}()
	f, err := os.Open(filename)
	defer closeFile(f)
	if err != nil {
		panic(err)
	}
	return f
}

func closeFile(f *os.File) {
	fmt.Println("close function working")
	f.Close()
}
