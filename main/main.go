package main

import (
	"fmt"
	"learning.com/awesomeProject/newmethods"
	"learning.com/awesomeProject/projects"
	"maps"
	"time"
)

func main() {
	fmt.Println(project.MyName())
	//value and variable
	v := "hello world"

	fmt.Println(v + ", welcome home")
	//constant
	const n = 5000000

	const t = 3e20 / n
	//for statement
	for {
		fmt.Println(t)
		break
	}

	for n := range 20 {
		if n%2 != 0 {
			continue
		}
		fmt.Println(n)
	}
	// switch statement
	weekday := time.Now()
	//weekday.Hour()
	switch weekday.Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")

	default:
		fmt.Println("It's a weekday")
		fmt.Println(weekday.Weekday())
		fmt.Println(weekday.Hour())
	}

	// array
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	b = [...]int{100, 3: 500, 600}
	fmt.Println(b)

	var r [2][3]int
	for i := range 2 {
		for j := range 3 {
			r[i][j] = i + j
		}
	}
	fmt.Println(r)

	r = [2][3]int{
		{2, 4, 5},
		{6, 8, 9},
	}
	fmt.Println("2d = ", r)

	//slice
	var s []string
	fmt.Println("emp =  ", s, "len =", len(s), "cap =", cap(s))

	s = make([]string, 3)
	fmt.Println("nil = ", s, "len =", len(s), "cap =", cap(s))
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)
	s = append(s, "d", "f", "g")
	fmt.Println("Appended slice is ", s)
	fmt.Println("len =", len(s), "cap =", cap(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy =", c)
	l3 := s[2:5]
	fmt.Println("sliced =", l3)
	fmt.Println("len =", len(l3), "cap =", cap(s))

	l := s[:5]
	fmt.Println("sliced =", l)
	fmt.Println("len =", len(l), "cap =", cap(s))

	l1 := s[2:]
	fmt.Println("sliced =", l1)
	fmt.Println("len =", len(l1), "cap =", cap(s))

	l4 := l1[2:]
	fmt.Println("sliced =", l4)
	fmt.Println("len =", len(l4), "cap =", cap(s))

	//maps
	m := make(map[int]string)
	m[1] = "robert"
	m[2] = "kevin"
	m[3] = "bob"

	fmt.Println(m)

	v1 := m[2]
	fmt.Println(v1)
	v2 := m[4]
	fmt.Println("v2 =", v2)

	fmt.Println("len =", len(m), "cap =", cap(s))

	delete(m, 2)
	fmt.Println(m)

	_, pr := m[1]
	fmt.Println(pr)

	clear(m)
	fmt.Println("m =")

	_, pr = m[2]
	fmt.Println(pr)

	f := map[string]int{"foo": 1, "bar": 2}
	f1 := map[string]int{"foo": 1, "bar": 2}

	if maps.Equal(f, f1) {
		fmt.Println("maps equal=======")
	}

	//functions

	sum := plus(3, 5)
	fmt.Println(sum)

	sum1 := add(6, 8, 40)
	fmt.Println(sum1)

	//multiple return value
	num, name := valu()

	nname, _ := valu()
	fmt.Println(num)
	fmt.Println(name)
	fmt.Println(nname)
	//variadic function "can be called with any number of trailing arguments"
	os := []int{1, 2, 3, 4, 5, 6}
	varfunc(os...)

	//closure "these are anonymous function that can be useful when you want to define a function inline without having to name it"
	useBcard1 := project.Activatecard()
	useBcard2 := project.Activatecard()
	useBcard3 := project.Activatecard()
	fmt.Println(useBcard1(200))
	fmt.Println(useBcard2(300))
	fmt.Println(useBcard3(400))

	//recursion
	fibb := project.Fib(7)
	fmt.Println(fibb)

	additi := project.Addi(3, 5)
	fmt.Println(additi, "====>")

	multiplication := project.Mult(8, 8)
	fmt.Println(multiplication)

	//struct
	c1 := project.NewCar("mazda", 4)
	fmt.Println(c1)

	//methods
	kl := project.Traipezium{Width: 10, Height: 5, Length: 6}
	fmt.Println(kl.Area(), "cm 2")

}

// multiple return value
func valu() (int, string) {
	return 34, "bob"
}

// functions
func plus(a int, b int) int {
	return a + b
}

func add(a, b, c int) int {
	return a + b + c
}

// variadic function
func varfunc(nums ...int) {
	fmt.Print(nums, " ==>")
	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum, "sum finished")

	cir := newmethods.Circle{Radius: 5}

	fmt.Println("the area of circle is :", cir.Area())

	//interface

	ro := newmethods.Round{Radius: 10}

	newmethods.Measure(ro)
}
