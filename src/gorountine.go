package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
 * `goroutine`은 Go 런타임이 관리하는 `논리적` 쓰레드다.
 * 함수 호출 시 `go` 키워드와 함께 호출하면 런타임시 새로운 `goroutine`을 실행한다.
 * 또한 `goroutine`은 비동기적으로 수행되기 때문에 여러 코드를 동시에 ( concurrency ) 실행하는데에 사용된다.
 * `goroutine`은 OS의 쓰레드와 1:1 대응되지 않고, 훨씬 적은 쓰레드를 사용한다.
 * 또한 `goroutine`은 기본적으로 1개의 `CPU`를 사용한다.
 */
func main() {
	runtime.GOMAXPROCS(8) // 8개의 Logical CPU 사용

	var wait sync.WaitGroup
	wait.Add(4)

	sayHi("sync", nil)
	go sayHi("Async 1", &wait)
	go sayHi("Async 2", &wait)
	//time.Sleep(time.Second * 1)

	go func() {
		defer wait.Done()
		fmt.Println("Async goroutine 1")
	}()

	go func() {
		defer wait.Done()
		fmt.Println("Async goroutine 2")
	}()
	wait.Wait()
}

func sayHi(s string, w *sync.WaitGroup) {
	if w != nil {
		defer w.Done()
	}

	for i:= 0; i < 5; i++ {
		fmt.Println(s, "Hi;", i + 1)
	}
}
