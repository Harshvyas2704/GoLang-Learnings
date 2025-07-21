package main

import (
	"fmt"
	"strings"
)

func main() {
	//var conferenceName = "Go Conference"
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	// var bookings [50]string -> Array
	// var bookings []string -> Slice
	bookings := []string{}

	// fmt.Println("Welcome to", conferenceName, "booking application")
	// fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	// fmt.Println("Thank you for booking in our conference")
	// fmt.Printf("Welcome to %v booking application\n", conferenceName)
	// fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)

	// data types
	// string interger boolean maps arrays

	// var bookings = [50]string{}

	for {
		var firstName string
		var lastName string
		var email string
		var useTickets uint

		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email address")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets you want to book")
		fmt.Scan(&useTickets)

		remainingTickets = remainingTickets - useTickets
		// bookings[0] = firstName + " " + lastName
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("User %v %v booked %v tickets. Tickets will be in your email %v\n", firstName, lastName, useTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
		// fmt.Printf("These is all the bookings, %v\n", bookings)
		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("User bookings %v\n", firstNames)
	}

	// fmt.Printf("ConferenceTickets is %T, RemainingTickets is %T\n", conferenceTickets, remainingTickets)

}
