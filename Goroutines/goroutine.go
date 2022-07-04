package main

import (
	"fmt"
	"sync"
	"time"
)

// GoRoutine --- >
func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func goroutine() {
	// direct call
	fun("direct call")

	// goroutine function call
	go fun("go-routine-1")

	// goroutine with anonymous function
	s := "goroutine-2"
	go func(s string) {
		fun(s)
	}(s)

	// goroutine with function value call
	fv := fun
	go fv("goroutine-3")

	time.Sleep(100 * time.Millisecond)
}

// WaitGroup --->
func waitGroup() {
	var wg sync.WaitGroup
	wg.Add(1)
	var data int
	go func() {
		wg.Done()
		data++
		data++
	}()
	wg.Wait()
	fmt.Println(data)
}

func closures() {
	var wa sync.WaitGroup
	incr := func(wg *sync.WaitGroup) {
		var i int
		wg.Add(1)
		go func() {
			defer wg.Done()
			i++
			fmt.Printf("value is %v", i)
		}()
		fmt.Println("return from function")
		return
	}
	incr(&wa)
	wa.Wait()
	fmt.Println("done....")
}

func main() {
	//goroutine()
	//waitGroup()
	closures()
}
