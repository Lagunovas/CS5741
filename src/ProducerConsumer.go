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
	"fmt"
	//"math/rand"
	//"strconv"
	"bufio"
	"os"
	"time"

	//linkedListCircularBuffer "github.com/CS5741/src/circularBuffer/nonConcurrent/linkedList"
	arrayCircularBuffer "github.com/CS5741/src/circularBuffer/nonConcurrent/array"
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
	numGenerator := NewNumberGenerator()
	buffer := arrayCircularBuffer.NewArayCircularBuffer(5)
	go Producer(buffer, numGenerator)
	go Consumer(buffer)

	reader := bufio.NewReader(os.Stdin)
	txt, _ := reader.ReadString('\n')
	fmt.Println(txt)
	fmt.Println("program complete")
}

/**/
func Producer(buffer *arrayCircularBuffer.ArrayCircularBuffer, numGen *NumberGenerator) {
	for i := 0; i < 20; i++ {
		num := numGen.GetNumber()
		if buffer.Size() == buffer.Capacity() {
			//sleep
			time.Sleep(2 * time.Second)
		}
		//add number to the bufer
		buffer.Push(num)
		fmt.Printf("Pushed number %d \n", num)
		if buffer.Size() == 1 {
			// wake up the consumer

		}

	}

}

/**/
func Consumer(buffer *arrayCircularBuffer.ArrayCircularBuffer) {
	for i := 0; i < 20; i++ {
		status, val := buffer.ReadNext()
		if status {
			fmt.Printf("consumer consumed %d \n", val)
		} else {
			time.Sleep(2 * time.Second)

		}

	}

}
