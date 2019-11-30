package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	concurrentArrayCircularBuffer "github.com/CS5741/src/circularBuffer/concurrent/array"
	concurrentLinkedListCircularBuffer "github.com/CS5741/src/circularBuffer/concurrent/linkedList"
	circularBufferInterface "github.com/CS5741/src/circularBuffer/interface"
	concurrentArrayStack "github.com/CS5741/src/stack/concurrent/array"
	concurrentBinaryTreeStack "github.com/CS5741/src/stack/concurrent/binaryTree"
	concurrentLinkedListStack "github.com/CS5741/src/stack/concurrent/linkedList"
	stackInterface "github.com/CS5741/src/stack/interface"
)

type RequesteeInterface interface {
	Request() (bool, int)
}

const BUFFER_SIZE int = 10

func main() {
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

	var dataStructure interface{}

	var results []time.Duration

	var startTime time.Time

	var nonParallelResults []time.Duration
	var parallelResults []time.Duration

	var resultsLocation *[]time.Duration

	for mode := 0; mode < 2; mode++ {
		switch mode {
		case 0:
			runtime.GOMAXPROCS(1)
			fmt.Println("Non-Parallel")
			resultsLocation = &nonParallelResults
		case 1:
			runtime.GOMAXPROCS(runtime.NumCPU())
			fmt.Printf("Parallel - THREAD COUNT: %v\n", runtime.GOMAXPROCS(0))
			resultsLocation = &parallelResults
		}

		for i := 2; i < 5; i++ {
			switch i {
			case 0:
				fmt.Println("===== Array Circular Buffer =====")
				dataStructure = concurrentArrayCircularBuffer.NewConcurrentArrayCircularBuffer(BUFFER_SIZE)
			case 1:
				fmt.Println("===== Linked List Circular Buffer =====")
				dataStructure = concurrentLinkedListCircularBuffer.NewConcurrentCircularBuffer(BUFFER_SIZE)
			case 2:
				fmt.Println("===== Array Stack =====")
				dataStructure = concurrentArrayStack.NewConcurrentArrayStack()
			case 3:
				fmt.Println("===== Binary Tree Stack =====")
				dataStructure = concurrentBinaryTreeStack.NewConcurrentBinaryTreeStack()
			case 4:
				fmt.Println("===== Linked List Stack =====")
				dataStructure = concurrentLinkedListStack.NewConcurrentLinkedListStack()
			}

			for j := 0; j < 10; j++ {
				startTime = time.Now()
				ProductionAndConsumption(numberOfProducers, numberOfConsumers, productionCapacity, consumptionCapacity, &waitGroup, &dataStructure)
				waitGroup.Wait()
				elapsedTime := time.Since(startTime)
				results = append(results, elapsedTime)
				fmt.Printf("time taken %v \n", results[len(results)-1])
			}

			resultCount := len(results)

			var total time.Duration
			duration, _ := time.ParseDuration("1h")
			var best time.Duration = duration
			duration, _ = time.ParseDuration("0ns")
			var worst time.Duration = duration

			for i := 0; i < resultCount; i++ {
				currentDuration := results[i]

				total += currentDuration

				if currentDuration < best {
					best = currentDuration
				}

				if currentDuration > worst {
					worst = currentDuration
				}
			}

			results = nil

			*resultsLocation = append(*resultsLocation, worst)
			*resultsLocation = append(*resultsLocation, best)
			*resultsLocation = append(*resultsLocation, total)

			fmt.Printf("Test results - worst: %v, best: %v, average: %vns, total: %v\n", worst, best.Nanoseconds(), total.Nanoseconds()/(int64)(resultCount), total)
		}
	}

	resultCount := len(nonParallelResults)

	for i := 0; i < resultCount; i++ {
		nPC := nonParallelResults[i]
		pC := parallelResults[i]

		fmt.Printf("Difference: %v\n", Difference(nPC.Nanoseconds(), pC.Nanoseconds()))
	}

}

const hundred float64 = 100

func Difference(v0, v1 int64) float64 {
	fv0 := float64(v0)
	fv1 := float64(v1)
	return -((fv1 - fv0) / fv0 * hundred)
}

func ProductionAndConsumption(numberOfProducers, numberOfConsumers, productionCapacityOfProducers, consumptionCapacityOfConsumer int, waitGroup *sync.WaitGroup, dataStructure *interface{}) {
	requestee := (*dataStructure).(RequesteeInterface)

	for i := 0; i < numberOfConsumers; i++ {
		waitGroup.Add(1)
		go Consumer(&requestee, consumptionCapacityOfConsumer, waitGroup)
	}

	if stack, ok := (*dataStructure).(stackInterface.StackInterface); ok {
		for i := 0; i < numberOfProducers; i++ {
			go StackProduction(i, &stack, productionCapacityOfProducers)
		}

	} else { //==========>BUFFER<================
		buffer, _ := (*dataStructure).(circularBufferInterface.CircularBufferInterface)

		for i := 0; i < numberOfProducers; i++ {
			go CircularBufferProduction(i, &buffer, productionCapacityOfProducers)
		}
	}

}

func StackProduction(id int, stack *stackInterface.StackInterface, productionCapacityOfProducers int) {
	for i := 1; i <= productionCapacityOfProducers; {
		num := (productionCapacityOfProducers * id) + i

		if (*stack).Size() < BUFFER_SIZE {
			(*stack).Push(num)
			i++
		}

		// fmt.Printf("Producer Produced %d \n", num)

		//sleep
		// time.Sleep(1 * time.Microsecond)
	}
}

func CircularBufferProduction(id int, buffer *circularBufferInterface.CircularBufferInterface, productionCapacityOfProducers int) {
	for i := 1; i <= productionCapacityOfProducers; {
		num := (productionCapacityOfProducers * id) + i

		val := (*buffer).Push(num)

		if val {
			// fmt.Printf("Producer Produced %d \n", num)
			i++
		}
	}

}

func Consumer(requestee *RequesteeInterface, consumptionCapacityOfConsumer int, waitGroup *sync.WaitGroup) {
	for i := 0; i < consumptionCapacityOfConsumer; {

		status, _ := (*requestee).Request()
		if status {
			// fmt.Printf("Consumer Consumed %d \n", val)
			i++
			// } else {
			// 	time.Sleep(1 * time.Millisecond)
			// fmt.Printf("FAIL - count: %v\n", i)
		}
	}

	waitGroup.Done()
}
