package project

func MyName() string {
	m := "robert karanja"

	return m
}

type car struct {
	name                     string
	doors                    int
	plateNumber              string
	driverName               string
	insuarenceExpirationDate int
}

func NewCar(name string, doors int) *car {
	c := car{name: name, doors: doors}
	c.plateNumber = "KCZ 180E"
	c.driverName = "Robert karanja"
	c.insuarenceExpirationDate = 2014
	return &c
}

// recursion fibonacci
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
func Addi(n, m int) int {
	if n == 0 {
		return m
	}
	return 1 + Addi(n-1, m)
}
func Mult(n, m int) int {
	if n == 0 {
		return 0
	}
	return m + Mult(n-1, m)
}

// closure
func Activatecard() func(int) int {
	amount := 1000
	debitfunc := func(debitamount int) int {
		amount -= debitamount
		return amount
	}
	return debitfunc
}

type Traipezium struct {
	Width, Length, Height float64
}

func (r Traipezium) Area() float64 {
	return 0.5 * (r.Width + r.Length) * r.Height
}
