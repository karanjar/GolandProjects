package interfaces

import (
	"fmt"
	"time"
)

func SendMessage(Msg Message) (string, int) {
	// ?
	Text := Msg.GetMessage()
	Cost := len(Text) * 3

	return Text, Cost
}

type Message interface {
	GetMessage() string
}

// don't edit below this line

type BirthdayMessage struct {
	BirthdayTime  time.Time
	RecipientName string
}

func (Bm BirthdayMessage) GetMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", Bm.RecipientName, Bm.BirthdayTime.Format(time.RFC3339))
}

type SendingReport struct {
	ReportName    string
	NumberOfSends int
}

func (Sr SendingReport) GetMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, Sr.ReportName, Sr.NumberOfSends)
}

type Employee interface {
	GetName() string
	GetSalary() int
}

type Contractor struct {
	Name         string
	HourlyPay    int
	HoursPerYear int
}

func (C Contractor) GetSalary() int {
	return C.HoursPerYear * C.HourlyPay
}

func (C Contractor) GetName() string {
	return C.Name
}

type FullTime struct {
	Name   string
	Salary int
}

func (Ft FullTime) GetSalary() int {
	return Ft.Salary
}

func (Ft FullTime) GetName() string {
	return Ft.Name
}

// multiple  interface
func (E Email) Cost() int {
	EmailLenth := len(E.Body)
	if E.IsSubscribed == true {
		return EmailLenth * 2
	} else {
		return EmailLenth * 5
	}

	// ?
}

func (E Email) Format() string {
	// ?
	if E.IsSubscribed == true {
		return fmt.Sprintf("'%s' | subscribed", E.Body)
	} else {
		return fmt.Sprintf("'%s' | Not subscribed", E.Body)
	}
}

type Expense interface {
	Cost() int
}

type Email struct {
	IsSubscribed bool
	Body         string
}

// formatter
type Formatter interface {
	Format() string
}

type Plaintext struct {
	Message string
}

func (Plp Plaintext) Format() string {
	return fmt.Sprintf("%s", Plp.Message)
}

type Bold struct {
	Message string
}

func (B Bold) Format() string {
	return fmt.Sprintf("**%s**", B.Message)
}

type Code struct {
	Message string
}

func (C Code) Format() string {
	return fmt.Sprintf("`%s`", C.Message)
}

func SendMessage1(Format Formatter) string {
	return Format.Format() // Adjusted to call Format without an argument
}
