package errors

import (
	"errors"
	"fmt"
)

func SendSMSToCouple(MsgToCustomer, MsgToSpouse string) (int, error) {
	// ?
	Cost1, err := SendSMS(MsgToCustomer)
	if err != nil {
		return 0, err
	}
	Cost2, err := SendSMS(MsgToSpouse)
	if err != nil {
		return 0, err
	}
	return Cost1 + Cost2, nil
}

// don't edit below this line

func SendSMS(Message string) (int, error) {
	const MaxTextLen = 25
	const CostPerChar = 2
	if len(Message) > MaxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", MaxTextLen)
	}
	return CostPerChar * len(Message), nil
}

func GetSMSErrorString(Cost float64, Recipient string) string {
	// ?

	return fmt.Sprintf("Send message to %s with cost $%.2f", Recipient, Cost)
}

type DivideError struct {
	Dividend float64
}

// ?
func (D DivideError) Error() string {
	return fmt.Sprintf("can't divide  %v by zero", D.Dividend)
}

func Divide(Dividend, Divisor float64) (float64, error) {
	if Divisor == 0 {
		return 0, DivideError{Dividend: Dividend}
	}
	return Dividend / Divisor, nil
}

func Divide1(X, Y float64) (float64, error) {
	if Y == 0 {
		// ?
		return 0, errors.New("can't divide by zero")
	}
	return X / Y, nil
}
func ValidateStatus(Status string) error {
	// ?
	if Status == "" {
		return errors.New("status can not be empty")
	}
	if len(Status) > 145 {
		return fmt.Errorf("status exceeds 140 characters ")
	}
	return nil
}
