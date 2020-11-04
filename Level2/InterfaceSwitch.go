package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Value: %v, type: %T (Integer).\n", v, v)
	case string:
		fmt.Printf("Value: %v, type: %T (String).\n", v, v)
	default:
		fmt.Printf("Value: %v, type: %T (Other type).\n", v, v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
