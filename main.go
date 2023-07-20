package main

import (
	"fmt"
	"strings"
	"sync"
	"booking-app/helper"
	"time"
)

var conferenceName = "GO Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = []string{}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketsNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
		if isValidName && isValidEmail && isValidTicketsNumber {

			bookTickets(remainingTickets, userTickets, firstName, lastName)
			wg.Add(1)
			go sendTicket(userTickets, firstName, email)

			firstName := getFirstName()
			fmt.Println("The first names of the bookings are: ", firstName)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Sorry, the conference is sold out.")
				// break
			}
		} else {
			if !isValidName {
				fmt.Println("Your name is invalid. Please try again.")
			}
			if !isValidEmail {
				fmt.Println("Your email is invalid. Please try again.")
			}
			if !isValidTicketsNumber {
				fmt.Println("Your ticket number is invalid. Please try again.")
			}
		}
		wg.Wait()
	}

}

func greetUsers() {

	fmt.Println("Welcome to ", conferenceName)
	fmt.Println("There are", conferenceTickets, "tickets available for the conference.")
	fmt.Println("We have", remainingTickets, "tickets left.")

}

func getFirstName() []string {

	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for name their name

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email adress:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets you would like to buy:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(remainingTickets uint, userTickets uint, firstName string, lastName string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %s %s, you have bought %d tickets for %s. \n", firstName, lastName, userTickets, conferenceName)
	fmt.Printf("They have %d tickets left. \n", remainingTickets)
}

func sendTicket(userTickets uint, fistName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("Hi %s, here is your ticket for %s. You have bought %d tickets. \n", fistName, conferenceName, userTickets)
	fmt.Println("##################")
	fmt.Printf("Sending ticket: ,%v \nto email address%v\n", email, ticket)
	fmt.Println("##################")
	wg.Done()
}
