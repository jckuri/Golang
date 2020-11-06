package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case quitNumber := <-quit:
			fmt.Println("quit", quitNumber)
			return
		}
	}
}

func PrintFibonacci(c, quit chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	quit <- 100
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go PrintFibonacci(c, quit)
	fibonacci(c, quit)
}

/*
OUTPUT:

0
1
1
2
3
5
8
13
21
34
quit 100
*/
