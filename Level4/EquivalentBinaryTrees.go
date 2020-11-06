package main

import (
    "golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int, level int) {
    if t.Left != nil {Walk(t.Left, ch, level + 1)}
	ch <- t.Value
	if t.Right != nil {Walk(t.Right, ch, level + 1)}	
	if level == 0 {close(ch)}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1, 0)
	go Walk(t2, ch2, 0)	
	for {
	    v1, ok1 := <- ch1
		v2, ok2 := <- ch2
		fmt.Printf("(%+v,%+v) (%+v,%+v)\n", v1, ok1, v2, ok2)
		if v1 != v2 {return false}
		if !ok1 || !ok2 {return ok1 == ok2}
	}
}

func main() {
    t1 := tree.New(1)
    t2 := tree.New(2)
	fmt.Println("Tree 1:", t1)
	fmt.Println("Tree 2:", t2)
	fmt.Printf("Same(tree1,tree2): %+v\n", Same(t1, t2))
	fmt.Printf("Same(tree1,tree1): %+v\n", Same(t1, t1))
}

/*
Tree 1: ((((1 (2)) 3 (4)) 5 ((6) 7 ((8) 9))) 10)
Tree 2: ((((2) 4 (6)) 8 (10 (12))) 14 ((16) 18 (20)))
(1,true) (2,true)
Same(tree1,tree2): false
(1,true) (1,true)
(2,true) (2,true)
(3,true) (3,true)
(4,true) (4,true)
(5,true) (5,true)
(6,true) (6,true)
(7,true) (7,true)
(8,true) (8,true)
(9,true) (9,true)
(10,true) (10,true)
(0,false) (0,false)
Same(tree1,tree1): true
*/
