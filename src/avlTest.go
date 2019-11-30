package main

import (
	"fmt"
	"math/rand"
	"sync"

	avl "github.com/CS5741/src/avl"
)

func Run() {
	//runtime.GOMAXPROCS(1)

	avlTree := avl.NewAVLTree()

	var wg sync.WaitGroup

	runs := 50

	var elements []int

	for i := 0; i < runs; i++ {
		val := (rand.Intn(1000) + 1)
		fmt.Println("Main: Starting worker", i)
		elements = append(elements, val)
		wg.Add(1)
		go putWorker(&wg, i, avlTree, val)
	}

	wg.Wait()

	avlTree.PrintTree(avlTree.Root)

	for i := 0; i < runs; i++ {
		wg.Add(1)
		go readerWorker(&wg, i, avlTree, elements[i])
	}

	wg.Wait()

	for i := 0; i < runs; i++ {
		wg.Add(1)
		go removerWorker(&wg, i, avlTree, elements[i])
	}

	fmt.Println("Main: Waiting for workers to finish")
	wg.Wait()
	fmt.Println("Main: Completed")

	avlTree.PrintTree(avlTree.Root)
}

func putWorker(wg *sync.WaitGroup, id int, avlTree *avl.AVLTree, value int) {
	defer wg.Done()
	status, returnValu := avlTree.Put(value)
	fmt.Printf("Put: %v, %v\n", status, returnValu)
}

func readerWorker(wg *sync.WaitGroup, id int, avlTree *avl.AVLTree, value int) {
	defer wg.Done()
	fmt.Printf("Get: %v\n", avlTree.Get(value))
}

func removerWorker(wg *sync.WaitGroup, id int, avlTree *avl.AVLTree, value int) {
	defer wg.Done()
	fmt.Printf("Remove: %v\n", avlTree.Remove(value))
}
