package main

import (
	"fmt"
	"math/rand"
	"sync"

	avl "github.com/CS5741/src/avl"
)

func main() {
	avlTree := avl.NewAVLTree()

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		fmt.Println("Main: Starting worker", i)
		wg.Add(1)
		go worker(&wg, i, avlTree, rand.Intn(10000))
	}

	fmt.Println("Main: Waiting for workers to finish")
	wg.Wait()
	fmt.Println("Main: Completed")

	avlTree.PrintTree(avlTree.Root)
}

func worker(wg *sync.WaitGroup, id int, avlTree *avl.AVLTree, value int) {
	defer wg.Done()

	fmt.Printf("Worker %v: Started\n", id)
	avlTree.Put(value)
	fmt.Printf("Worker %v: Finished\n", id)
}
