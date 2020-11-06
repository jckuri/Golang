/*
Statement
Create a channel of int and then create a goroutine to add a value to the channel and then print the channel value in the main function
 
Topics to Practice: 
goroutine, channel, function
*/

package main

import (
	"fmt"
	"time"
)

func goroutine(ch chan int) {
	for i:=0; i<10; i++ {
		ch <- i
		time.Sleep(500 * time.Millisecond)
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go goroutine(ch)
	for value := range ch {
		fmt.Printf("value: %v (%v)\n", value, time.Now())
	}
}

/*
value: 0 (2009-11-10 23:00:00 +0000 UTC m=+0.000000001)
value: 1 (2009-11-10 23:00:00.5 +0000 UTC m=+0.500000001)
value: 2 (2009-11-10 23:00:01 +0000 UTC m=+1.000000001)
value: 3 (2009-11-10 23:00:01.5 +0000 UTC m=+1.500000001)
value: 4 (2009-11-10 23:00:02 +0000 UTC m=+2.000000001)
value: 5 (2009-11-10 23:00:02.5 +0000 UTC m=+2.500000001)
value: 6 (2009-11-10 23:00:03 +0000 UTC m=+3.000000001)
value: 7 (2009-11-10 23:00:03.5 +0000 UTC m=+3.500000001)
value: 8 (2009-11-10 23:00:04 +0000 UTC m=+4.000000001)
value: 9 (2009-11-10 23:00:04.5 +0000 UTC m=+4.500000001)
*/
