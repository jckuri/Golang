/*
Statement
Create a goroutine that will execute an anonymous function to print “Hello World” and in the main routine print “main function”

Topics to Practice: 
goroutine, function, common Issue
*/

package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%v %v\n", s, i)
	}
}

func main() {
	go say("Hello World")
	say("main function")
}

/*
OUTPUT:

main function 0
Hello World 0
Hello World 1
main function 1
main function 2
Hello World 2
Hello World 3
main function 3
main function 4
*/
