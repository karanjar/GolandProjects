package channels

func Sendt(Ch chan int) {
	for i := 0; i < 5; i++ {
		Ch <- i
	}

	close(Ch)
}
