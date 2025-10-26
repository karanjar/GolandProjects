package goroutines

import (
	"fmt"
	"sync"
)

var (
	Counters int
	Wg       sync.WaitGroup
	Couter   int
	Mu       sync.Mutex
	RW       sync.RWMutex
)

func Increment(Wg *sync.WaitGroup) {

	defer Wg.Done()
	Mu.Lock()
	Couter++
	Mu.Unlock()

}
func Readcounter() {
	defer Wg.Done()
	RW.RLock()
	fmt.Println("Reading counters", Counters)
	RW.RUnlock()
}
func Writecounter() {
	defer Wg.Done()
	RW.Lock()
	Counters++
	RW.Unlock()
}
