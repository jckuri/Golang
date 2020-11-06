/*
Given a list of strings find all strings that contain a given substring. 
You should split the list into chunks and process them in parallel. 
Each goroutine should write the response to an output channel which will be consumed in the end to consolidate the response.

strs := []string{"abc", "bcd", "efg", "aabcr", "acb", "ggg", "hjuklbc"}
substr := "bc"
*/

package main

import (
	"fmt"
	"strings"
	"sync"
)

func IntMin(a int, b int) int {
	if a < b {return a}
	return b
}

func ProcessChunk(strs []string, substr string, channel chan string, wg *sync.WaitGroup) {
	fmt.Printf("Chunk: %v, substr: %v\n", strs, substr)
	for _, str := range strs {
		if strings.Contains(str, substr) {
			channel <- str
		}
	}
	wg.Done()
}

func ProcessChunks(strs []string, substr string, n_chunks int, channel chan string, wg *sync.WaitGroup) {
	n := len(strs)
	if n == 0 {
		fmt.Errorf("You need at least 1 string.\n")
	}
	if n_chunks > n {
		fmt.Errorf("The number of chunks cannot be greater than the number of strings.\n")
	}
	chunk_inc := int(float64(n) / float64(n_chunks) + 0.5)
	fmt.Printf("chunk_inc=%v, n_chunks=%v\n", chunk_inc, n_chunks)
	for i:=0; i<n; i+=chunk_inc {
		i1:=i
		i2:=IntMin(i+chunk_inc, n)
		fmt.Printf("i1=%v, i2=%v\n", i1, i2)
		wg.Add(1)
		go ProcessChunk(strs[i1:i2], substr, channel, wg)
	}
	wg.Wait()
	close(channel)
}

func main() {
	strs := []string{"abc", "bcd", "efg", "aabcr", "acb", "ggg", "hjuklbc"}
	substr := "bc"
	fmt.Printf("strs=%v\n", strs)
	channel := make(chan string)
	var wg sync.WaitGroup
	go ProcessChunks(strs, substr, 2, channel, &wg)
	for s := range channel {
		fmt.Printf("Received: %v\n", s)
	}
}

/*
./prog.go:36:13: result of fmt.Errorf call not used
./prog.go:39:13: result of fmt.Errorf call not used
Go vet exited.

strs=[abc bcd efg aabcr acb ggg hjuklbc]
chunk_inc=4, n_chunks=2
i1=0, i2=4
i1=4, i2=7
Chunk: [acb ggg hjuklbc], substr: bc
Received: hjuklbc
Chunk: [abc bcd efg aabcr], substr: bc
Received: abc
Received: bcd
Received: aabcr

Program exited.
*/
