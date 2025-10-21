package goroutines

import "sync"

var (
	Couter int
	Mu     sync.Mutex
)

func Increment(Wg *sync.WaitGroup) {
	defer Wg.Done()
	Couter++

}
