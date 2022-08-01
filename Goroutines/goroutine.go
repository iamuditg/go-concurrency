package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"sync/atomic"
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
	//closures()
	//practice()
	//debitCredit()
	//syncChannel()

	var wgt sync.WaitGroup
	var mutx sync.Mutex
	wgt.Add(2)
	go func() {
		defer wgt.Done()
		for i := 0; i < 10; i++ {
			mutx.Lock()
			fmt.Println("Hello")
			mutx.Unlock()
			time.Sleep(1000 * time.Millisecond)
		}
	}()
	go func() {
		defer wgt.Done()
		for i := 0; i < 10; i++ {
			mutx.Lock()
			fmt.Println("Udit")
			mutx.Unlock()
			time.Sleep(1000 * time.Millisecond)
		}
	}()
	wgt.Wait()
}

func syncChannel() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go func(ch1 <-chan string, ch2 chan<- string) {
			fmt.Println(<-ch1)
			ch2 <- "Pong"
			wg.Done()
		}(ch1, ch2)
		go func(ch2 <-chan string) {
			defer wg.Done()
			fmt.Println(<-ch2)
		}(ch2)
		ch1 <- "Ping"
		wg.Wait()
	}
}

var mutex = &sync.Mutex{}
var balance int64

// sync call in go routine
func debitCredit() {
	var wg sync.WaitGroup
	balance = 200
	fmt.Println("Intital balance is ", balance)
	wg.Add(1)
	go credit(&wg)
	wg.Add(1)
	go debit(&wg)
	wg.Wait()
	fmt.Println("final balance is ", balance)
}

func debit(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		mutex.Lock()
		atomic.AddInt64(&balance, -100)
		fmt.Println("After debiting balance is ", balance)
		mutex.Unlock()
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func credit(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		mutex.Lock()
		atomic.AddInt64(&balance, 100)
		fmt.Println("After crediting balance is ", balance)
		mutex.Unlock()
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

var wg sync.WaitGroup
var goRoutine sync.WaitGroup

func practice() {

	//chn := make(chan int)
	//wg.Add(2)
	//go responseSize("http://stackoverflow.com", chn)
	//go responseSize("http://google.com", chn)
	//fmt.Println(<-chn)
	//wg.Wait()
	//close(chn)
	//time.Sleep(10 * time.Second)

	//wg.Add(1)
	//command := make(chan string)
	//go routine(command, &wg)
	//time.Sleep(1 * time.Second)
	//command <- "Pause"
	//time.Sleep(1 * time.Second)
	//command <- "Play"
	//time.Sleep(1 * time.Second)
	//command <- "Stop"
	//wg.Wait()

	//wg.Add(3)
	//go increment("Python")
	//go increment("Java")
	//go increment("Golang")
	//wg.Wait()
	//fmt.Println("counter: ", counter)

	//rand.Seed(time.Now().Unix())
	//project := make(chan string, 10)
	//goRoutine.Add(5)
	//for i := 0; i <= 5; i++ {
	//	go employee(project, i)
	//}
	//for j := 1; j <= 10; j++ {
	//	project <- fmt.Sprintf("Project: %d", i)
	//}
	//close(project)
	//goRoutine.Wait()

	//var val int
	//var wg sync.WaitGroup
	////ch := make(chan int)
	//var mut sync.Mutex
	//wg.Add(2)
	//mut.Lock()
	//go func() {
	//	mut.Unlock()
	//	val = 1
	//	//ch <- val
	//	wg.Done()
	//}()
	//go func() {
	//	mut.Lock()
	//	fmt.Println(val)
	//	wg.Done()
	//}()
	//wg.Wait()

}

func employee(projects chan string, employee int) {
	defer goRoutine.Done()
	for {
		project, result := <-projects
		if result == false {
			fmt.Printf("Employee: %d : Exit \n", employee)
			return
		}
		fmt.Printf("Employee: %d Started %s \n", employee, project)

		// Random wait to simulate work time
		sleep := rand.Int63n(50)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Println("\n TIme of sleep", sleep, "ms \n")
		fmt.Printf("Employee: %d completed %s \n", employee, project)
	}
}

var (
	counter int32
	wgh     sync.WaitGroup
	mut     sync.Mutex
)

func increment(name string) {
	defer wg.Done()
	for range name {
		//atomic.AddInt32(&counter, 1)
		//runtime.Gosched()
		mut.Lock()
		counter++
		mut.Unlock()
	}
}

var i int

func work() {
	time.Sleep(250 * time.Millisecond)
	i++
	fmt.Println(i)
}

func routine(command <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	var status = "Play"
	for {
		select {
		case cmd := <-command:
			fmt.Println(cmd)
			switch cmd {
			case "Stop":
				return
			case "Pause":
				status = "Pause"
			default:
				status = "Play"
			}
		default:
			if status == "Play" {
				work()
			}
		}
	}
}

func responseSize(url string, chn chan int) {
	defer wg.Done()
	fmt.Println("Step1: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Step2: ", len(body))
	chn <- len(body)
}
