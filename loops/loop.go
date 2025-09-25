package loops

import "fmt"

func Bulksend(NumMessage int) float64 {
	Totalcost := 0.0

	for i := 0; i <= NumMessage; i++ {
		fee := float64(i) * 0.01
		cost := 1.0 + fee
		Totalcost += cost
	}
	return Totalcost
}

func FizzBuzz() {
	// ?
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println("Fizz")
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
		}
		if i%3 == 0 || i%5 == 0 {
			fmt.Println("FizzBuzz")
		}
	}

}
