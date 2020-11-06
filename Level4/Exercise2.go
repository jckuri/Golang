/*
Statement
Create a function goroutine that will execute an anonymous function to just print the number “1”, in the main function print the number “0” and also add a time.Sleep() to wait 2 seconds.
 
Topics to Practice: 
goroutine, function, time pkg
*/

package main

import (
	"fmt"
	"time"
)

func goroutine() {
	go func() {
		fmt.Println("1")
	}()
}

func main() {
	goroutine()
	fmt.Println("0")
	time.Sleep(2 * time.Second)
}

/*
OUTPUT:

0
1
*/
