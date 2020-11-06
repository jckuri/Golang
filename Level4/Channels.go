package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int, ms time.Duration) {
	sum := 0
	for _, v := range s {
		sum += v
		time.Sleep(ms * time.Millisecond)
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	n2 := len(s)/2
	s1 := s[:n2]
	s2 := s[n2:]

	fmt.Println(s1, s2)

	c := make(chan int)
	
	go sum(s1, c, 1000)
	go sum(s2, c, 1)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

/*
[7 2 8] [-9 4 0]
-5 17 12
*/
