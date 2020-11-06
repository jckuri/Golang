package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	response := make(chan *http.Response, 1)
	errors := make(chan *error)

	go func() {
		resp, err := http.Get("http://matt.aimonetti.net/")
		if err != nil {
			errors <- &err
		}
		response <- resp
	}()
	for {
		select {
		case r := <-response:
			fmt.Printf("BODY:\n%s\n", r.Body)
			return
		case err := <-errors:
			log.Fatal(*err)
		case <-time.After(2000 * time.Millisecond):
			fmt.Printf("Timed out!")
			return
		}
	}
}

/*
OUTPUT:

$ go run Timeout.go 
BODY:
&{[] {%!s(*http.http2clientStream=&{0xc000082600 0xc00019a000 <nil> 1 0xc000090480 {{0 0} {{} <nil> {0 0 0 <nil> <nil>} 0} 0xc000068100 0 <nil> <nil> <nil> <nil>} false true 0x645a80 {[] 65536 0xc000082660} {[] 4194304 0xc000082670} -1 <nil> <nil> false 0xc000024fc0 <nil> 0xc000025020 true true false 0 map[] 0xc000168198})} %!s(*gzip.Reader=<nil>) <nil>}
*/
