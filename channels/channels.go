package main

import "fmt"

func channels() {
	ch := make(chan int)
	go func(a int, b int) {
		c := a + b
		ch <- c
	}(1, 2)
	fmt.Println(<-ch)
}

func main() {
	//channels()
	//rangeCh()
	//bufferedChannel()
	//channelDirection()
	//channelOwnerShip()
}

func channelOwnerShip() {
	owner := func() <-chan int {
		ch := make(chan int)
		go func() {
			defer close(ch)
			for i := 0; i < 5; i++ {
				ch <- i
			}
		}()
		return ch
	}
	consumer := func(ch <-chan int) {
		for v := range ch {
			fmt.Printf("Received %d\n", v)
		}
		fmt.Println("Done Receiving!")
	}
	ch := owner()
	consumer(ch)
}

func channelDirection() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go genMsg(ch1)
	go recvMsg(ch1, ch2)
	rec := <-ch2
	fmt.Println(rec)

}

func genMsg(ch1 chan<- string) {
	// send Message to ch1
	ch1 <- "message"
}

func recvMsg(ch1 <-chan string, ch2 chan<- string) {
	// recv message on ch1
	// sent it on ch2
	m := <-ch1
	ch2 <- m
}

func bufferedChannel() {
	ch := make(chan int, 6)
	go func() {
		defer close(ch)
		for i := 0; i < 6; i++ {
			fmt.Printf("Sending %d\n", i)
			ch <- i
		}
	}()

	for v := range ch {
		fmt.Printf("Received %v\n", v)
	}
}

func rangeCh() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println("sent")
			ch <- i
		}
		close(ch)
	}()
	for v := range ch {
		fmt.Println(v)
	}
}
