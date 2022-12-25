package main

import (
	"booking/helper"
	"fmt"
	"sync"
	"time"
)

const confernceTicket = 50

var confernceName = "go Conference!"
var remainingTicket uint = 50
var bookings = make([]userData, 0)

type userData struct {
	firstName  string
	lastName   string
	email      string
	noOfTicket uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	// for {
	firstName, lastName, email, noOfTicket := userinput()
	isValidName, isValidEmail, isValidTicket := helper.ValidateUserInput(firstName, lastName, email, noOfTicket, remainingTicket)

	if isValidName && isValidEmail && isValidTicket {
		bookTicket(noOfTicket, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(noOfTicket, firstName, lastName, email)

		firstNames := getFirstName()
		fmt.Printf("all firstname of  booking least : %v \n", firstNames)

		if remainingTicket == 0 {
			fmt.Println("all the ceets are full come to next confrence ")
			// break
		}
	} else {
		if !isValidName {
			fmt.Println("name is not Valid")
		}
		if !isValidEmail {
			fmt.Println("enter the valid email id ")
		}
		if !isValidTicket {
			fmt.Println("ticket is not valid")

		}
	}
	// }
	wg.Wait()
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", confernceName)

	fmt.Printf("we have totel of =%v avilabel and remaining = %v so Book fast...!!\n ", confernceTicket, remainingTicket)
	fmt.Println("get yout ticket hear to attend the confrence")
}

func getFirstName() []string {
	firstNames := []string{}
	//  "_ " = index
	for _, booking := range bookings {

		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func userinput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var noOfTicket uint
	fmt.Println("Enter your firft name ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email ")
	fmt.Scan(&email)
	fmt.Println("Enter number of Ticket ")
	fmt.Scan(&noOfTicket)
	return firstName, lastName, email, noOfTicket
}

func bookTicket(noOfTicket uint, firstName string, lastName string, email string) {
	remainingTicket = remainingTicket - noOfTicket

	//map creation
	var userData = userData{
		firstName:  firstName,
		lastName:   lastName,
		email:      email,
		noOfTicket: noOfTicket,
	}

	bookings = append(bookings, userData)
	fmt.Printf("list of booking %v \n", bookings)
	fmt.Println("----------------------------------")
	fmt.Println("booking Detials -> ")
	fmt.Printf("name = %v %v \n Email = %v \n No Of Ticket = %v \n", firstName, lastName, email, noOfTicket)
	fmt.Println("----------------------------------")
	fmt.Println("THANKE YOU...!")
	fmt.Printf("remaining ticket = %v avalible for this confurnce %v \n", remainingTicket, confernceName)

}

func sendTicket(noOfTicket uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v ticket for %v %v ", noOfTicket, firstName, lastName)
	fmt.Println("################")
	fmt.Printf("sinding ticket:\n %v \nto email %v\n", ticket, email)
	fmt.Println("################")
	wg.Done()
}
