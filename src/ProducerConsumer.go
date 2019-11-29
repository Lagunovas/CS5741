// package producerConsumer

// import(
// 	"fmt"
// )

// type Producer struct{
// 	id string
// }
package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"time"

	//=======> BUFFER
	concurrentArrayCircularBuffer "github.com/CS5741/src/circularBuffer/concurrent/array"
	circularBufferInterface "github.com/CS5741/src/circularBuffer/interface"
	//concurrentBinaryTreeCircularBuffer "github.com/CS5741/src/circularBuffer/concurrent/binaryTree"
	//concurrentLinkedListCircularBuffer "github.com/CS5741/src/circularBuffer/concurrent/linkedList"

	//=========> Stack
	concurrentArrayStack "github.com/CS5741/src/stack/concurrent/array"
	//concurrentBinaryTreeStack "github.com/CS5741/src/stack/concurrent/binaryTree"
	//concurrentLinkedListStack "github.com/CS5741/src/stack/concurrent/linkedList"
	stackInterface "github.com/CS5741/src/stack/interface"
)

const STACK = 1
const BUFFER = 2

//
type NumberGenerator struct {
	number int
}

type Producer struct {
	id                 int
	productionCapacity int
}
type Consumer struct {
	id                  int
	consumptionCapacity int
}

//
func NewNumberGenerator() *NumberGenerator {
	return &NumberGenerator{number: 0}
}

func (numGen *NumberGenerator) GetNumber() int {
	numGen.number++
	return numGen.number
}

func main() {

	runtime.GOMAXPROCS(1)
	//fmt.Println(runtime.GOMAXPROCS(0))
	ProductionAndConsumption(2, 2, 100, 1, 100)

	reader := bufio.NewReader(os.Stdin)

	txt, _ := reader.ReadString('\n')
	fmt.Println(txt)
	fmt.Println("program complete")
}

/**/

/**/

func ProductionAndConsumption(numberOfProducers, numberOfConsumers, productionCapacityOfProducers, typeOfProduction, consumptionCapacityOfConsumer int) {

	start := time.Now()
	if typeOfProduction == STACK {
		var stack stackInterface.StackInterface = concurrentArrayStack.NewConcurrentArrayStack()
		//	var stack stackInterface.StackInterface = concurrentBinaryTreeStack.NewConcurrentBinaryTreeStack()
		//var stack stackInterface.StackInterface = concurrentLinkedListStack.NewConcurrentLinkedListStack()
		for i := 0; i < numberOfProducers; i++ {
			numGenerator := NewNumberGenerator()
			go StackProduction(stack, numGenerator, productionCapacityOfProducers)
		}
		for i := 0; i < numberOfConsumers; i++ {
			go StackConsumption(start, stack, consumptionCapacityOfConsumer)
		}
	} else {
		var buffer circularBufferInterface.CircularBufferInterface = concurrentArrayCircularBuffer.NewConcurrentArrayCircularBuffer(5)
		// var buffer circularBufferInterface.CircularBufferInterface = concurrentBinaryTreeCircularBuffer.NewConcurrentBinaryTreeCircularBuffer(5)
		//var buffer circularBufferInterface.CircularBufferInterface = concurrentLinkedListCircularBuffer.NewConcurrentCircularBuffer(5)
		for i := 0; i < numberOfProducers; i++ {
			numGenerator := NewNumberGenerator()
			go CircularBufferProduction(buffer, numGenerator, productionCapacityOfProducers)
		}
		for i := 0; i < numberOfConsumers; i++ {
			go CircularBufferConsumption(start, buffer, consumptionCapacityOfConsumer)
		}
	}

}

func StackProduction(stack stackInterface.StackInterface, numGen *NumberGenerator, productionCapacityOfProducers int) {
	for i := 0; i < productionCapacityOfProducers; i++ {
		num := numGen.GetNumber()
		stack.Push(num)
		fmt.Printf("Producer Produced %d \n", num)
		//sleep
		time.Sleep(1 * time.Millisecond)
	}
}

func CircularBufferProduction(buffer circularBufferInterface.CircularBufferInterface, numGen *NumberGenerator, productionCapacityOfProducers int) {
	for i := 0; i < productionCapacityOfProducers; i++ {
		num := numGen.GetNumber()
		for buffer.Size() == buffer.Capacity() {
			//sleep
			time.Sleep(1 * time.Millisecond)
		}
		val := buffer.Push(num)
		if val {
			fmt.Printf("Producer Produced %d \n", num)
		} else {
			fmt.Printf("could not push number %d \n", num)
			i--
		}
	}
}

func StackConsumption(startTime time.Time, stack stackInterface.StackInterface, consumptionCapacityOfConsumer int) {
	for i := 0; i < consumptionCapacityOfConsumer; i++ {
		status, val := stack.Pop()
		if status {
			fmt.Printf("consumer consumed %d \n", val)
		} else {
			time.Sleep(1 * time.Millisecond)
			fmt.Println("consumer failed")
			i--
		}
		if i == consumptionCapacityOfConsumer-1 {
			elapsed := time.Since(startTime)
			fmt.Printf("time taken %s", elapsed)
		}

	}
}
func CircularBufferConsumption(startTime time.Time, buffer circularBufferInterface.CircularBufferInterface, consumptionCapacityOfConsumer int) {
	for i := 0; i < consumptionCapacityOfConsumer; i++ {
		status, val := buffer.ReadNext()
		if status {
			fmt.Printf("consumer consumed %d \n", val)
		} else {
			time.Sleep(1 * time.Millisecond)
			fmt.Println("consumer failed")
			i--
		}
		if i == consumptionCapacityOfConsumer-1 {
			elapsed := time.Since(startTime)
			fmt.Printf("time taken %s", elapsed)
		}
	}
}
