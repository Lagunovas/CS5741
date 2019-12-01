package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"

	avl "github.com/CS5741/src/avl"
)

func Run() {
	runtime.GOMAXPROCS(1)

	avlTree := avl.NewAVLTree()

	var wg sync.WaitGroup

	runs := 5

	var elements []int
	fmt.Println("===========>Put<===========")
	for i := 0; i < runs; i++ {
		val := (rand.Intn(1000) + 1)
		fmt.Println("Main: Starting worker", i)
		elements = append(elements, val)
		wg.Add(1)
		go putWorker(&wg, i, avlTree, val, i)
	}
	wg.Wait()
	fmt.Println("===========>Put<===========")

	fmt.Println("========>TREE<===========")
	avlTree.PrintTree(avlTree.Root)
	fmt.Println("========>TREE<===========")

	fmt.Println("========>Get<===========")
	for i := 0; i < runs; i++ {
		wg.Add(1)
		fmt.Println("Main: Starting worker", i)
		go readerWorker(&wg, i, avlTree, elements[i], i)
	}
	wg.Wait()
	fmt.Println("========>Get<===========")

	fmt.Println("========>Remove<===========")
	for i := 0; i < runs; i++ {
		wg.Add(1)
		fmt.Println("Main: Starting worker", i)
		go removerWorker(&wg, i, avlTree, elements[i], i)
	}
	fmt.Println("Main: Waiting for workers to finish")
	wg.Wait()
	fmt.Println("========>Remove<===========")

	fmt.Println("========>TREE<===========")

	avlTree.PrintTree(avlTree.Root)
	fmt.Println("========>TREE<===========")

}

func putWorker(wg *sync.WaitGroup, id int, avlTree *avl.AVLTree, value int, worker int) {
	defer wg.Done()
	status, returnValu := avlTree.Put(value, worker)
	fmt.Printf("worker %d status = %v, value = %v\n", worker, status, returnValu)
}

func readerWorker(wg *sync.WaitGroup, id int, avlTree *avl.AVLTree, value, worker int) {
	defer wg.Done()
	fmt.Printf("Get worker: %d value = %v\n", worker, avlTree.Get(value, worker))
}

func removerWorker(wg *sync.WaitGroup, id int, avlTree *avl.AVLTree, value, worker int) {
	defer wg.Done()
	fmt.Printf("Remove worker: %d value = %v\n", worker, avlTree.Remove(value, worker))
}
