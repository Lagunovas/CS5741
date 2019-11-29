// package producerConsumer

// import(
// 	"fmt"
// )

// type Producer struct{
// 	id string
// }
package main

import (
	"fmt"
	"sync"
	"time"

	//=======> BUFFER
	concurrentArrayCircularBuffer "github.com/CS5741/src/circularBuffer/concurrent/array"
	circularBufferInterface "github.com/CS5741/src/circularBuffer/interface"

	//concurrentBinaryTreeCircularBuffer "github.com/CS5741/src/circularBuffer/concurrent/binaryTree"

	//=========> Stack
	// concurrentArrayStack "github.com/CS5741/src/stack/concurrent/array"
	//concurrentBinaryTreeStack "github.com/CS5741/src/stack/concurrent/binaryTree"
	//concurrentLinkedListStack "github.com/CS5741/src/stack/concurrent/linkedList"
	stackInterface "github.com/CS5741/src/stack/interface"
)

type RequesteeInterface interface {
	Request() (bool, int)
}

func main() {
	//runtime.GOMAXPROCS(1)
	var waitGroup sync.WaitGroup
	var numberOfProducers int
	var numberOfConsumers int
	var consumptionCapacity int
	var productionCapacity int

	fmt.Println("Please Enter the numberOfProducers : ")
	fmt.Scanln(&numberOfProducers)
	fmt.Println("Please Enter the numberOfConsumers : ")
	fmt.Scanln(&numberOfConsumers)
	fmt.Println("Please Enter the consumptionCapacity : ")
	fmt.Scanln(&consumptionCapacity)
	fmt.Println("Please Enter the productionCapacity : ")
	fmt.Scanln(&productionCapacity)

	// var arrayStack stackInterface.StackInterface = concurrentArrayStack.NewConcurrentArrayStack()
	// var binaryTreeStack stackInterface.StackInterface = concurrentBinaryTreeStack.NewConcurrentBinaryTreeStack()
	// var linkedListStack stackInterface.StackInterface = concurrentLinkedListStack.NewConcurrentLinkedListStack()

	var arrayBuffer circularBufferInterface.CircularBufferInterface = concurrentArrayCircularBuffer.NewConcurrentArrayCircularBuffer(5)
	// var binaryTreeBuffer circularBufferInterface.CircularBufferInterface = concurrentBinaryTreeCircularBuffer.NewConcurrentBinaryTreeCircularBuffer(5)
	//var linkedListBuffer circularBufferInterface.CircularBufferInterface = concurrentLinkedListCircularBuffer.NewConcurrentCircularBuffer(5)
	//fmt.Println(runtime.GOMAXPROCS(0))

	startTime := time.Now()

	for i := 0; i < 1; i++ {
		ProductionAndConsumption(numberOfProducers, numberOfConsumers, productionCapacity, consumptionCapacity, &waitGroup, arrayBuffer)
		waitGroup.Wait()
	}

	elapsedTime := time.Since(startTime)
	fmt.Printf("time taken %s \n", elapsedTime)
}

func ProductionAndConsumption(numberOfProducers, numberOfConsumers, productionCapacityOfProducers, consumptionCapacityOfConsumer int, waitGroup *sync.WaitGroup, datastructure interface{}) {
	requestee := datastructure.(RequesteeInterface)

	if stack, ok := datastructure.(stackInterface.StackInterface); ok {
		for i := 0; i < numberOfProducers; i++ {
			go StackProduction(i, stack, productionCapacityOfProducers)
		}

		for i := 0; i < numberOfConsumers; i++ {
			waitGroup.Add(1)
			go Consumer(requestee, consumptionCapacityOfConsumer, waitGroup)
		}

	} else { //==========>BUFFER<================
		buffer, _ := datastructure.(circularBufferInterface.CircularBufferInterface)

		for i := 0; i < numberOfProducers; i++ {
			go CircularBufferProduction(i, buffer, productionCapacityOfProducers)
		}

		for i := 0; i < numberOfConsumers; i++ {
			waitGroup.Add(1)
			go Consumer(requestee, consumptionCapacityOfConsumer, waitGroup)
		}

	}

}

func StackProduction(id int, stack stackInterface.StackInterface, productionCapacityOfProducers int) {
	for i := 1; i <= productionCapacityOfProducers; i++ {
		num := (productionCapacityOfProducers * id) + i

		stack.Push(num)
		fmt.Printf("Producer Produced %d \n", num)

		//sleep
		time.Sleep(1 * time.Millisecond)
	}
}

func CircularBufferProduction(id int, buffer circularBufferInterface.CircularBufferInterface, productionCapacityOfProducers int) {
	for i := 1; i <= productionCapacityOfProducers; {
		num := (productionCapacityOfProducers * id) + i

		for buffer.Size() == buffer.Capacity() {
			time.Sleep(1 * time.Millisecond)
		}

		val := buffer.Push(num)

		if val {
			// fmt.Printf("Producer Produced %d \n", num)
			i++
		}
	}

}

func Consumer(requestee RequesteeInterface, consumptionCapacityOfConsumer int, waitGroup *sync.WaitGroup) {
	for i := 0; i < consumptionCapacityOfConsumer; {
		status, _ := requestee.Request()
		if status {
			//fmt.Printf("Consumer Consumed %d \n", val)
			i++
		} else {
			time.Sleep(1 * time.Millisecond)
			// fmt.Printf("FAIL - count: %v\n", i)
		}
	}

	waitGroup.Done()
}
