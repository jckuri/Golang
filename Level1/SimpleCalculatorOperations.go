/*
Simple Calculator Operations
Statement
Create a calculator struct with a result attribute and with the most common methods: 
Add
Subtract
Multiply
Divide


Plus: Handle possible error, Divide function won’t be able to handle second parameter with value zero, instead function should panic and print an alert

Topics to Practice:
function, return value, panic, defer
*/

package main

import (
	"fmt"
	"runtime"
	"reflect"
)

type Calculator struct {
	a, b, result float64
}

func Add(c *Calculator) {
	c.result = c.a + c.b
}

func Subtract(c *Calculator) {
	c.result = c.a - c.b
}

func Multiply(c *Calculator) {
	c.result = c.a * c.b
}

func Divide(c *Calculator) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("ERROR: Division by zero (%v).\n", r)
		}
	}()
	c.result = c.a / c.b
	if c.b == 0 {panic(1)}	
}

func ComputeOperation(operation Operation) {
	operation.operation(operation.calculator)
}

type Operation struct {
	calculator *Calculator
	operation func(c *Calculator)
}

// Taken from: https://stackoverflow.com/questions/7052693/how-to-get-the-name-of-a-function-in-go
func GetFunctionName(i interface{}) string {
    return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func (o Operation) String() string {
	return fmt.Sprintf("%v(%v,%v)=%v", GetFunctionName(o.operation), o.calculator.a, o.calculator.b, o.calculator.result)
}

func main() {
	operations := [5]Operation {{&Calculator{a: 1, b: 2}, Add}, {&Calculator{a: 2, b: 3}, Multiply}, {&Calculator{a: 5, b: 2}, Subtract}, {&Calculator{a: 12, b: 0}, Divide}, {&Calculator{a: 6, b: 3}, Divide}}
	for _, operation := range operations {
		ComputeOperation(operation)
		fmt.Println(operation)
	}
}

/*
OUTPUT:

main.Add(1,2)=3
main.Multiply(2,3)=6
main.Subtract(5,2)=3
ERROR: Division by zero (1).
main.Divide(12,0)=+Inf
main.Divide(6,3)=2
*/
