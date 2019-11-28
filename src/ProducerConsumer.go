// package producerConsumer

// import(
// 	"fmt"
// )

// type Producer struct{
// 	id string
// }
package main

import (
	//"bytes"
	"bufio"
	"fmt"
	"os"
	"runtime"
	"time"

	//"math/rand"
	//"strconv"

	linkedListCircularBuffer "github.com/CS5741/src/circularBuffer/nonConcurrent/linkedList"
	//arrayCircularBuffer "github.com/CS5741/src/circularBuffer/nonConcurrent/array"
)

//
type NumberGenerator struct {
	number int
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
	start := time.Now()
	runtime.GOMAXPROCS(1)
	fmt.Println(runtime.GOMAXPROCS(0))
	numGenerator := NewNumberGenerator()
	buffer := linkedListCircularBuffer.NewLinkedListCircularBuffer(5)
	//buffer := arrayCircularBuffer.NewArayCircularBuffer(5)
	go Producer(buffer, numGenerator)
	go Consumer(start, buffer)

	reader := bufio.NewReader(os.Stdin)

	txt, _ := reader.ReadString('\n')
	fmt.Println(txt)
	fmt.Println("program complete")
}

/**/
func Producer(buffer *linkedListCircularBuffer.LinkedListCircularBuffer, numGen *NumberGenerator) {
	for i := 0; i < 100; i++ {
		num := numGen.GetNumber()

		for buffer.Size() == buffer.Capacity() {
			//sleep
			time.Sleep(1 * time.Millisecond)
		}
		val := buffer.Push(num)
		if val {
			fmt.Printf("Producer Produced %d \n", num)
		} else {
			fmt.Print("could not push number %d \n", num)
			i--
		}
		//add number to the bufer
		if buffer.Size() == 1 {
			// wake up the consumer
		}
	}
}

/**/
func Consumer(startTime time.Time, buffer *linkedListCircularBuffer.LinkedListCircularBuffer) {
	for i := 0; i < 100; i++ {
		status, val := buffer.ReadNext()
		if status {
			fmt.Printf("consumer consumed %d \n", val)
		} else {
			time.Sleep(1 * time.Millisecond)
			fmt.Println("consumer failed")
			i--
		}
		if i == 99 {
			elapsed := time.Since(startTime)
			fmt.Printf("time taken %s", elapsed)
		}
	}
}
