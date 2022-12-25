package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, noOfTicket uint, remainingTicket uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := noOfTicket > 0 && noOfTicket <= remainingTicket
	return isValidName, isValidEmail, isValidTicket
}
