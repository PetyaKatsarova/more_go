package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

func Walk(t *tree.Tree, ch chan int) {
	if t == nil { return }
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 0; i < 10; i++ {
		val1, val2 := <-ch1, <-ch2
		if val1 != val2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch) //creates sorted binary tree where each node has val of multiple of K: in this case 1

	fmt.Println("First 10 values from the channel:")
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("R tree.New(1) and tree.New(1) the same tree?", Same(tree.New(1), tree.New(1))) // true
	fmt.Println("R tree.New(1) and tree.New(2) the same tree?", Same(tree.New(1), tree.New(2))) // false
}
/*
tree.New(1) creates a binary tree where the root node has the value 1, the left subtree contains nodes with values 2, 3, 4, and 
so on (multiples of 1), and the right subtree contains nodes with values 0, -1, -2, and so on (negatives of multiples of 1).
tree.New(2) creates a binary tree where the root node has the value 2, the left subtree contains nodes with values 4, 6, 8, 
and so on (multiples of 2), and the right subtree contains nodes with values 0, -2, -4, and so on (negatives of multiples of 2).
*/