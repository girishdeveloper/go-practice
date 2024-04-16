package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Program Start")
	var waitGroup sync.WaitGroup
	waitGroup.Add(10)
	for i := 0; i < 10; i++ {
		go concurrentTasks(i, &waitGroup)
	}
	waitGroup.Wait()
	finishTask()
	fmt.Println("Program End")
}

func finishTask() {
	fmt.Println("All tasks are finished")
}

func concurrentTasks(taskNo int, waitGroup *sync.WaitGroup) {
	fmt.Printf("Start task # %d\n", taskNo)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("End task # %d\n", taskNo)
	waitGroup.Done()
}
