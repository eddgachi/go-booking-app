package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

// var conferenceName string = "Go conference"
var conferenceName = "Go conference"

const conferenceTickets int = 50

var remainingTickets uint = 50

// var bookings = [50]string{}
// var bookings []string // changed to slice
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()

		fmt.Printf("The first names of bookings are %v\n", firstNames)

		noTicketsRemaining := remainingTickets == 0
		if noTicketsRemaining {
			fmt.Println("Our conference is booked out. Come back next year.")
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
	wg.Wait()

}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)

	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
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

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// bookings[0] = firstName + " " + lastName // add items to an array
	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	var ticket = fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName)
	time.Sleep(10 * time.Second)
	fmt.Println("################")
	fmt.Printf("Sending Ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("################")
	wg.Done()
}
