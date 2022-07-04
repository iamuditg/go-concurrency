package main

import (
	"fmt"
	"time"
)

func goSelect() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-ch1:
			fmt.Println(m1)
		case m2 := <-ch2:
			fmt.Println(m2)
		}
	}
}

func main() {
	//goSelect()
	//timeoutSelect()
	nonBlockingCommunication()
}

func nonBlockingCommunication() {
	ch := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(1 * time.Second)
			ch <- "message"
		}
	}()

	for i := 0; i < 2; i++ {
		select {
		case m := <-ch:
			fmt.Println(m)
		default:
			fmt.Println("no message receive")
		}

		fmt.Println("processing..")
		time.Sleep(1500 * time.Millisecond)

	}
}

func timeoutSelect() {
	ch := make(chan string, 1)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- "one"
	}()
	select {
	case m := <-ch:
		fmt.Println(m)
	case <-time.After(4 * time.Second):
		fmt.Println("timeout")
	}
}
