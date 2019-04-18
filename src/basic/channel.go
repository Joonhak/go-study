package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

/*
 * Go의 Channel: 데이터를 주고받는 통로
 */
func main() {
	// 사용할 `논리적` CPU 설정
	runtime.GOMAXPROCS(8)

	ch := make(chan int, 2)

	go func() { // 익명함수를 통한 `channel`에 data 전달
		ch <- 12345
	}()
	go fmt.Println(<-ch) // 12345

	go sendToChannel(ch, 123)
	go receiveFromChannel(ch) // 123

	boolCh := make(chan bool, 2)
	count := 5

	go func() {
		for i := 0; i < count; i++ {
			boolCh <- true
			fmt.Println("Go routine: ", i+1)
		}
		boolCh <- true
	}()

	for i := 0; i < count; i++ {
		<-boolCh // 해당 채널이 종료될때 까지 대기
		fmt.Println("Main function: ", i+1)
	}

	ch2 := make(chan int, 2)

	ch <- 654
	ch <- 321

	close(ch2)

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	if _, success := <-ch2; !success {
		fmt.Println("No more data")
	}

	strCh := make(chan string, 2)
	go channelRange(strCh)

	receiveAfterClosed(strCh)

	channelSelect()

}

func sendToChannel(ch chan<- int, data int) { // 송신 전용 channel
	ch <- data
}

func receiveFromChannel(ch <-chan int) { // 수신 전용 channel
	fmt.Println(<-ch)
}

func channelRange(ch chan<- string) {
	for i := 0; i < 5; i++ {
		ch <- "Send string data! " + strconv.Itoa(i+1)
	}

	close(ch)
}

func receiveAfterClosed(ch <-chan string) {
	for data := range ch { // range 문을 통해 `channel`이 종료되기 전까지 수신한다.
		fmt.Println(data)
	}
}

func channelSelect() {
	done1 := make(chan bool, 2)
	done2 := make(chan bool, 2)

	go sleepOneSec(done1)
	go sleepTwoSec(done2)

EXIT:
	for {
		select {
		case <-done1:
			fmt.Println("Channel 1 done.")

		case <-done2:
			fmt.Println("Channel 2 done.")
			break EXIT
		}
	}

}

func sleepOneSec(ch chan bool) {
	time.Sleep(time.Second)
	ch <- true
}
func sleepTwoSec(ch chan bool) {
	time.Sleep(time.Second * 2)
	ch <- true
}
