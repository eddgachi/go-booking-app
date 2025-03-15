package main

import (
	"fmt"
	"strings"
)

// var conferenceName string = "Go conference"
var conferenceName = "Go conference"

const conferenceTickets int = 50

var remainingTickets uint = 50

// var bookings = [50]string{}
// var bookings []string // changed to slice
var bookings = []string{}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames(bookings)

			fmt.Printf("The first names of bookings are %v\n", firstNames)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}

			// alternative way of doing the conditional
			// if remainingTickets == 0 {
			// 	fmt.Println("Our conference is booked out. Come back next year.")
			// 	break
			// }
		} else {
			if !isValidName {
				fmt.Println("the first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("the email address you entered doesn't contain an @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("the number of tickets you entered is invalid")
			}
		}

	}

}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)

	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}

	for _, booking := range bookings {
		var names = strings.Fields(booking)
		var firstName = names[0]
		firstNames = append(firstNames, firstName)
	}

	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	// bookings[0] = firstName + " " + lastName // add items to an array
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
