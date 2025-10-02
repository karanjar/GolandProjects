package main

import (
	"fmt"
	"maps"
	"time"

	"learning.com/awesomeProject/errors"
	"learning.com/awesomeProject/interfaces"
	"learning.com/awesomeProject/loops"
	"learning.com/awesomeProject/newmethods"
	project "learning.com/awesomeProject/projects"
	"learning.com/awesomeProject/slices"
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
	fmt.Println("FOR STATEMENT")
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
	fmt.Println("SWITCH STATEMENT")
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
	fmt.Println("ARRAY=====")
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

	getmes, cost3 := slices.GetMessageWithRetries("jfvjjvjfj", "eufuieuiefu", "eifiiieei")
	fmt.Println(getmes, "the cost is ", cost3)

	//slice
	fmt.Print("SLICE=====")
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

	messG, err := slices.GetMessageWithRetriesForPlan("proo", [3]string{"Hello", "you look great today", "have a nice weekend"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(messG)

	messcost := slices.GetMessageCosts([]string{"Hello", "you look great today", "have a nice weekend", "or can we share some info", "The weekend is over", "Have a nice weekday"})
	fmt.Println(messcost)
	c0st := slices.Sum(2, 4, 6, 7, 9, 0, 4, 3, 3)
	fmt.Println(c0st)

	dayg := []slices.Cost{
		{Day: 1, Value: 10.2},
		{Day: 2, Value: 10.5},
		{Day: 3, Value: 10.6},
		{Day: 4, Value: 10.7},
		{Day: 5, Value: 10.8},
		{Day: 6, Value: 10.9},
	}

	dayc := slices.GetDayCosts(dayg, 4)
	fmt.Println(dayc)

	badword := []string{"fuck", "mother fucker", "father fucker", "bitch"}
	meswaord := []string{"", "food", "ugali", "sugar"}

	indexof := slices.IndexOfFirstBadWord(meswaord, badword)
	fmt.Println(indexof)
	d2 := slices.CreateMatrix(10, 10)
	for _, row := range d2 {
		for _, col := range row {
			fmt.Printf("%4d", col)
		}
		fmt.Println("=======================================")
	}

	//maps
	fmt.Println("MAPS======")
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
	fmt.Println("FUNCTIONS=======")
	sum := plus(3, 5)
	fmt.Println(sum)

	sum1 := add(6, 8, 40)
	fmt.Println(sum1)

	//multiple return value
	fmt.Println("MULTIPLE RETURN=====")
	num, name := valu()

	nname, _ := valu()
	fmt.Println(num)
	fmt.Println(name)
	fmt.Println(nname)
	//variadic function "can be called with any number of trailing arguments"
	os := []int{1, 2, 3, 4, 5, 6}
	varfunc(os...)

	//closure "these are anonymous function that can be useful when you want to define a function inline without having to name it"
	fmt.Println("CLOSURE=========")
	useBcard1 := project.Activatecard()
	useBcard2 := project.Activatecard()
	useBcard3 := project.Activatecard()
	fmt.Println(useBcard1(200))
	fmt.Println(useBcard2(300))
	fmt.Println(useBcard3(400))

	//recursion
	fmt.Println("RECURSION=========")
	fibb := project.Fib(7)
	fmt.Println(fibb)

	additi := project.Addi(3, 5)
	fmt.Println(additi)

	multiplication := project.Mult(8, 8)
	fmt.Println(multiplication)

	//struct
	fmt.Println("STRUCT==========")
	c1 := project.NewCar("mazda", 4)
	fmt.Println(c1)

	//methods
	fmt.Println("METHOD=========")
	kl := project.Traipezium{Width: 10, Height: 5, Length: 6}
	fmt.Println(kl.Area(), "cm 2")

	//interface
	fmt.Println("INTERFACES==========")
	ro := newmethods.Round{Radius: 10}

	newmethods.Measure(ro)

	hbd := interfaces.BirthdayMessage{BirthdayTime: time.Date(2004, 8, 23, 22, 45, 67, 0, time.UTC),
		RecipientName: "ROBERT KARANJA"}

	mr := interfaces.SendingReport{ReportName: "Monthly report",
		NumberOfSends: 60}
	text1, cost1 := interfaces.SendMessage(hbd)
	text2, cost2 := interfaces.SendMessage(mr)

	fmt.Println(text1, cost1)
	fmt.Println(text2, cost2)

	ctr := interfaces.Contractor{Name: "peterson warotere",
		HourlyPay:    150,
		HoursPerYear: 3600}

	ft := interfaces.FullTime{Name: "Benson Karanja"}
	fmt.Println(ctr)
	fmt.Println(ctr.GetSalary())
	fmt.Println(ft)

	email := interfaces.Email{
		IsSubscribed: true,
		Body:         "Hello World",
	}

	fmt.Println(email.Format())

	////enums
	fmt.Println("ENUMS=======")
	//news := newmethods.Transform(newmethods.Carready)
	//fmt.Println(news)
	//
	//car := news
	//fmt.Println(car)

	//struct embedding
	fmt.Println("STRUCT EMBEDDING========")
	tx := user{
		name:   "bob",
		number: "0707712786",
		sender: sender{
			rateLimit: 50,
		},
	}
	fmt.Println(tx.rateLimit)
	//struct methods

	auth := authenticationInfo{
		username: "bob",
		password: "kldfvjfvj",
	}
	fmt.Println(auth.grtBasicAuth())

	customer := newUser("bob", "standard")
	fmt.Println(customer)
	customerMessage, _ := customer.sendMassage("i what do you want to do", 1000)
	fmt.Println("message sent", customerMessage)

	plaintext := interfaces.Plaintext{
		Message: "KARATINA UNIVERSTY STUDENTS",
	}
	bold := interfaces.Bold{
		Message: "KARATINA UNIVERSTY STUDENTS",
	}
	code := interfaces.Code{
		Message: "KARATINA UNIVERSTY STUDENTS",
	}
	fmt.Println(interfaces.SendMessage1(plaintext))
	fmt.Println(interfaces.SendMessage1(bold))
	fmt.Println(interfaces.SendMessage1(code))

	//error
	totalcost, err := errors.SendSMSToCouple("The message has", "Your spouse has send ")
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	fmt.Println(totalcost)
	smsE := errors.GetSMSErrorString(34.4535, "ROBERT KARANJA")
	fmt.Println(smsE)

	divderC, err := errors.Divide(3, 0)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println(divderC)
	}
	divs, err := errors.Divide1(3, 6)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	fmt.Println(divs)
	err = errors.ValidateStatus("countyu in  country")
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	tc := loops.Bulksend(5)
	fmt.Println(tc)

	loops.FizzBuzz()

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

}

// struct embedding
type sender struct {
	rateLimit int
}

type user struct {
	name   string
	number string
	sender
}

//struct methods in go

type authenticationInfo struct {
	username string
	password string
}

// create the method below
func (a authenticationInfo) grtBasicAuth() string {
	return fmt.Sprintf("Authorization: Basic %s:%s", a.username, a.password)
}

func (U User) sendMassage(message string, messagelenth int) (string, bool) {

	if messagelenth <= U.MessagecharLimit {
		return message, true
	} else {
		return "", false
	}

}

type User struct {
	Name string
	Membership
}
type Membership struct {
	Type             string
	MessagecharLimit int
}

func newUser(name string, membershipType string) User {
	// ?
	limit := 1000
	if membershipType == "premium" {
		limit = 1000

	} else {
		limit = 100
	}
	return User{
		Name: name,
		Membership: Membership{
			Type:             membershipType,
			MessagecharLimit: limit,
		},
	}
}
