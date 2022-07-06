package main

import (
	"fmt"
	"sync"
)

func main() {
	//pipelines()
	fanInFanOut()
}

func fanInFanOut() {
	in := generator(2, 3)
	ch1 := square(in)
	ch2 := square(in)
	for n := range merge(ch1, ch2) {
		fmt.Println(n)
	}

}

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, cs := range cs {
		go output(cs)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// pipelines
func pipelines() {
	for out := range square(square(generator(2, 3))) {
		fmt.Println(out)
	}
}

func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}