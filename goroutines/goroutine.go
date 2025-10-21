package goroutines

import (
	"fmt"

	"time"
)

func Printnumbers() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Am the other goroutine...")

}
