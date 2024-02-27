package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("conferenceName is %T, remainingTickets is %T, conferenceTickets is %T\n", conferenceName, remainingTickets, conferenceTickets)

	fmt.Printf("Welcome to %v Booking Application\n", conferenceName)
	fmt.Printf("We Have total of %v Tickets and %v are still remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get Your Ticket here to attend")

	var userName string
	var userTickets int
	//ask user for their name

	userName = "Brian"
	userTickets = 50
	fmt.Printf("User %v Booked %v Tickets\n", userName, userTickets)

}
