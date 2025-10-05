package slices

import (
	"errors"
)

func GetMessageWithRetries(Primary, Secondary, Tertiary string) ([3]string, [3]int) {

	Message := [3]string{Primary, Secondary, Tertiary}

	costs := [3]int{20, 12, 30}

	return Message, costs

}

const (
	PlanFree = "free"
	PlanPro  = "pro"
)

func GetMessageWithRetriesForPlan(Plan string, Messages [3]string) ([]string, error) {

	if Plan == PlanPro {
		return Messages[:3], nil
	}
	if Plan == PlanFree {
		return Messages[:2], nil
	}
	return nil, errors.New("unsupported plan")

}

func GetMessageCosts(Messages []string) []float64 {
	Messcost := make([]float64, len(Messages))
	for I, Message := range Messages {
		Messcost[I] = float64(len(Message)) * 0.01
	}

	return Messcost
}
func Sum(Nums ...int) int {
	C0st := 0
	for I := 0; I < len(Nums); I++ {
		C0st += Nums[I]
	}
	return C0st
}

type Cost struct {
	Day   int
	Value float64
}

func GetDayCosts(Costs []Cost, Day int) []float64 {
	var Result []float64
	for _, C := range Costs {
		if C.Day == Day {
			Result = append(Result, C.Value)
		}
	}
	return Result

}

func IndexOfFirstBadWord(Msg []string, BadWords []string) int {
	for I, Word := range Msg {
		for _, BadWord := range BadWords {
			if Word == BadWord {
				return I
			}
		}
	}
	return -1
}
func CreateMatrix(Rows, Cols int) [][]int {
	// ?
	var Data [][]int
	for I := 1; I < Rows; I++ {
		Row := []int{}
		for J := 1; J < Cols; J++ {
			Row = append(Row, I*J)
		}
		Data = append(Data, Row)
	}
	return Data

}

type Message interface {
	Type() string
}

type TextMessage struct {
	Sender  string
	Content string
}

func (tm TextMessage) Type() string {
	return "text"
}

type MediaMessage struct {
	Sender    string
	MediaType string
	Content   string
}

func (mm MediaMessage) Type() string {
	return "media"
}

type LinkMessage struct {
	Sender  string
	URL     string
	Content string
}

func (lm LinkMessage) Type() string {
	return "link"
}

// Don't touch above this line

func FilterMessages(Messages []Message, FilterType string) []Message {
	// ?
	var MessageF []Message
	for _, M := range Messages {
		if FilterType == M.Type() {
			MessageF = append(MessageF, M)
		}

	}
	return MessageF

}

func IsValidPassword(Password string) bool {
	// ?
	if len(Password) < 8 {
		return false
	}
	Hasupper := false
	Hasdigit := false
	for I := 0; I < len(Password); I++ {
		char := Password[I]

		if char >= '0' && char <= '9' {
			Hasdigit = true
		}
		if char >= 'A' && char <= 'Z' {
			Hasupper = true
		}

	}
	return Hasdigit && Hasupper
}
