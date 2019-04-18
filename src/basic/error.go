package main

import (
	"log"
	"os"
)

type myErr struct {
	msg string
}

// Implements error interface
func (m myErr) Error() string {
	return m.msg
}

func makeErr() (string, error) {
	x := myErr{"some error occurred!"}
	return "", x
}

func main() {
	f, err := os.Open("D:\\backup\\기본적인 정보.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Print(f.Name())

	_, err = makeErr()

	// `custom error`에 대한 처리
	switch err.(type) {
	default: // no error

	case myErr:
		log.Println("My Error occurred:", err.Error())
	case error:
		log.Fatal(err.Error())
	}

}
