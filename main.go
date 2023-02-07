package main

import (
	//	"booking-app/helper"
	"fmt"
	//	"strings"
	//	"strconv"
)

var conferenceName = "GO Conference"

const conferenceTickets = 50

var remainingTickets = 50

// var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets int
}

func main() {

	greetUsers()
	for {
		// get User Input
		firstName, lastName, email, userTicket := getUserInput()

		// Validate user input
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTicket, remainingTickets)
		if isValidName && isValidEmail && isValidTicketNumber {
			// book ticket
			bookTicket(firstName, lastName, email, userTicket)

			// Calling function print first Names
			//	printFirstNames(bookings)
			firstNames := printFirstNames()
			fmt.Printf("These are all of our booking:%v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Sorry we are sold out")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("Either your first name or last name is too short")

			}
			if !isValidEmail {
				fmt.Println("Your Email is not valid")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets is not valid")
			}

		}

	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and  %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get Your tickets to attend")
}

func printFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// var names = strings.Fields(booking)
		// var firstName = names[0]
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames

}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTicket int
	fmt.Println("Please enter your First name")
	fmt.Scan(&firstName)
	fmt.Println("Please enter your Last name")
	fmt.Scan(&lastName)
	fmt.Println("Please enter your Email address")
	fmt.Scan(&email)
	fmt.Println("Please enter the number of tickets you want to book:")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket
}

func bookTicket(firstName string, lastName string, email string, userTicket int) {
	remainingTickets = remainingTickets - userTicket
	//creating structs

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTicket,
	}
	bookings = append(bookings, userData)
	fmt.Printf("List of Bookings:%v\n", bookings)
	// create a map of users
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["noOfTickets"] = strconv.FormatInt(int64(userTicket), 10)
	// bookings = append(bookings, userData)
	fmt.Printf("List of Bookings:%v\n", bookings)

	fmt.Printf("The Whole Slice:%v\n", bookings)
	fmt.Printf("The first value:%v\n", bookings[0])
	fmt.Printf("Slice type:%T\n", bookings)
	fmt.Printf("Slice length:%v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a conformation email at %v\n", firstName, lastName, userTicket, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
