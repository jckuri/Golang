/*
Statement
Create a Stack structure that means a LIFO (Last In First Out).
Note: The func main should be provided and Just need to create the Stack struct and Methods to Push, Pop and Peek. IsEmpty is optional.
To test it, push two or three strings and then pop the last one and finally, push another one.
*/

package main
import (
	"fmt"
)

type Stack struct {
    elements []string
}

func NewStack() Stack {
    return Stack{make([]string, 0)}
}

func (stack *Stack) Push(s string) {
    stack.elements = append(stack.elements, s)
}

func (stack *Stack) Pop() string {
    n := len(stack.elements)
    element := stack.elements[n - 1]
    stack.elements = stack.elements[:n - 1]
    return element
}

func (stack *Stack) Peek() string {
    n := len(stack.elements)
    element := stack.elements[n - 1]
    return element
}

func (stack *Stack) IsEmpty() bool {
    n := len(stack.elements)
    return n == 0
}

func main() {
    stack := NewStack()
    fmt.Println("IsEmpty: ", stack.IsEmpty())
    fmt.Println(stack)
    stack.Push("Hello")
    fmt.Println(stack)
    stack.Push("World")
    fmt.Println(stack)
    element := stack.Pop()
    fmt.Println(element, stack)
    stack.Push("Goodbye")
    fmt.Println(stack)
    element = stack.Peek()
    fmt.Println(element, stack)
    fmt.Println("IsEmpty: ", stack.IsEmpty())
}
