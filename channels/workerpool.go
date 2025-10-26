package channels

import (
	"fmt"
	"sync"
	"time"
)

func WorkerPool(Id int, Tasks <-chan int, Results chan<- int, WG *sync.WaitGroup) {
	defer WG.Done()
	for Task := range Tasks {
		fmt.Printf("Task %d is picked up by the worker %d\n", Task, Id)
		time.Sleep(1 * time.Second)
		Results <- Task * 2
	}

}
