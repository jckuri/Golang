/*
Statement
Create a function that will increase a number and that function will be executed by a goroutine inside a for loop (x1000 times). To avoid race conditioning, implement the sync.Mutex and Lock and Unlock inside the increase() function. 
Note: Add a time.Sleep() to be able to see the final n
 
Topics to Practice: 
goroutine, sync.mutex, function, for loop, defer, time pkg
*/

package main

import (
	"fmt"
	"time"
	"sync"
)

type SharedCounter struct {
	counter int
	mutex sync.Mutex
}

func (sc *SharedCounter) Inc1() {
	sc.mutex.Lock()
	sc.counter++
	sc.mutex.Unlock()
}

func (sc *SharedCounter) Counter() int {
	sc.mutex.Lock()
	defer sc.mutex.Unlock()
	return sc.counter
}

func main() {
	var sc SharedCounter
	for i:=0; i<1000; i++ {
		go func() {
			sc.Inc1()
		} ()
	}
	time.Sleep(time.Second)
	fmt.Printf("Counter: %v\n", sc.Counter())
}

/*
OUTPUT:

Counter: 1000
*/
