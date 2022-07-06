package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
	atomic2 "sync/atomic"
	"time"
)

func main() {
	//mutex()
	//atomic()
	//conditionalVar()
	//syncOnce()
	syncPool()
}

var bufPool = sync.Pool{New: func() interface{} {
	fmt.Println("allocating new bytes.Buffer")
	return new(bytes.Buffer)
}}

func log(w io.Writer, debug string) {
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	b.WriteString(time.Now().Format("15:04:05"))
	b.WriteString(":")
	b.WriteString(debug)
	b.WriteString("\n")
	w.Write(b.Bytes())
	bufPool.Put(b)
}

func syncPool() {
	log(os.Stdout, "debug-string1")
	log(os.Stdout, "debug-string2")
}

func syncOnce() {
	var wg sync.WaitGroup
	var once sync.Once
	load := func() {
		fmt.Println("Run only once intialization function")
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
		}()
		once.Do(load)
	}
	wg.Wait()
}

func mutex() {
	var balance int
	var wg sync.WaitGroup
	var mu sync.Mutex
	deposit := func(amount int) {
		mu.Lock()
		balance += amount
		mu.Unlock()
	}
	withdrawn := func(amount int) {
		mu.Lock()
		defer mu.Unlock()
		balance -= amount
	}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			deposit(1)
		}()
	}
	wg.Wait()

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			withdrawn(1)
		}()
	}
	wg.Wait()
	fmt.Println(balance)
}

var sharedRsc = make(map[string]interface{})

func conditionalVar() {
	var wg sync.WaitGroup
	mu := sync.Mutex{}
	c := sync.NewCond(&mu)
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.L.Lock()
		for len(sharedRsc) == 0 {
			c.Wait()
		}
		fmt.Println(sharedRsc["rsc1"])
		c.L.Unlock()
	}()
	c.L.Lock()
	sharedRsc["rsc1"] = "hello"
	c.Signal()
	c.L.Unlock()
	c.Wait()
}

func atomic() {
	runtime.GOMAXPROCS(4)
	var counter uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for c := 0; c < 1000; c++ {
				atomic2.AddUint64(&counter, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Printf("counter: %v\n", counter)
}
