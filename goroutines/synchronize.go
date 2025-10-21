package goroutines

import (
	"fmt"
	"sync"
	"time"
)

func Worler(i int, delay int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker No. %v processing started\n", i)

	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Printf("Worker No. %v processing finished\n", i)
}
