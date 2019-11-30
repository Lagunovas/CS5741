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
	"runtime"
	"sync"
	"time"

	//=======> BUFFER
	concurrentArrayCircularBuffer "github.com/CS5741/src/circularBuffer/concurrent/array"
	circularBufferInterface "github.com/CS5741/src/circularBuffer/interface"

	concurrentLinkedListCircularBuffer "github.com/CS5741/src/circularBuffer/concurrent/linkedList"

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
	// runtime.GOMAXPROCS(1)
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

	var buffer circularBufferInterface.CircularBufferInterface
	//var linkedListBuffer circularBufferInterface.CircularBufferInterface = concurrentLinkedListCircularBuffer.NewConcurrentCircularBuffer(5)
	//fmt.Println(runtime.GOMAXPROCS(0))

	var results []time.Duration

	var startTime time.Time

	for mode := 0; mode < 2; mode++ {
		switch mode {
		case 0:
			runtime.GOMAXPROCS(1)
			fmt.Println("Non-Parallel")
		case 1:
			runtime.GOMAXPROCS(12)
			fmt.Printf("Parallel - THREAD COUNT: %v\n", runtime.GOMAXPROCS(0))
		}

		for i := 0; i < 3; i++ {
			switch i {
			case 0:
				buffer = concurrentArrayCircularBuffer.NewConcurrentArrayCircularBuffer(5)
			case 1:
				buffer = concurrentLinkedListCircularBuffer.NewConcurrentCircularBuffer(5)
			}

			for j := 0; j < 10; j++ {
				startTime = time.Now()
				ProductionAndConsumption(numberOfProducers, numberOfConsumers, productionCapacity, consumptionCapacity, &waitGroup, buffer)
				waitGroup.Wait()
				elapsedTime := time.Since(startTime)
				results = append(results, elapsedTime)
				fmt.Printf("time taken %s \n", results[len(results)-1])
			}
		}
	}

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
			// time.Sleep(1 * time.Millisecond)
			// fmt.Printf("FAIL - count: %v\n", i)
		}
	}

	waitGroup.Done()
}
