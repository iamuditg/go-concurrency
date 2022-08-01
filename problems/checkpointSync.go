package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

/*
The checkpoint synchronization is a problem of synchronizing multiple tasks.
Consider a workshop where several workers assembling details of some mechanism.
When each of them completes his work, they put the details together.
There is no store, so a worker who finished its part first must wait for others
before starting another one.
Putting details together is the checkpoint at which tasks synchronize themselves
before going their paths apart.
*/

var (
	partList    = []string{"A", "B", "C", "D"}
	nAssemblies = 3
	wg          sync.WaitGroup
)

func worker(part string) {
	log.Println(part, "worker begins part")
	time.Sleep(time.Duration(rand.Int63n(1e6)))
	log.Println(part, "worker completes part")
	wg.Done()
}
