package channels

import (
	"fmt"
	"sync"
)

func Workers(i int, Ch chan int) {
	fmt.Println("sending message to channel woker...")
	Ch <- i
	fmt.Println("sent message to channel  woker!")
}
func Send(Ch chan<- int, Wg *sync.WaitGroup) {
	defer Wg.Done()
	fmt.Println("sending message to a receive only channel...")
	Ch <- 20
}
func Receive(Ch <-chan int, Wg *sync.WaitGroup) {
	defer Wg.Done()

	fmt.Println("waiting message from a send only channel...")
	Msg := <-Ch
	fmt.Println("received message from a receive only channel", Msg)
}
