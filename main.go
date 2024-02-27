package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v Booking Application\n", conferenceName)
	fmt.Printf("We Have total of %v Tickets and %v are still remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get Your Ticket here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int
	//ask user for their name
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address:")
	fmt.Scan(&email)
	fmt.Println("Enter Number of tickets you need:")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("User %v %v with email %v Booked %v Tickets\n", firstName, lastName, email, userTickets)
	fmt.Printf("Remaining tickets are: %v\n", remainingTickets)

}
