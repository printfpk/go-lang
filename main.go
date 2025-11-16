package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	fmt.Printf("welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have total %v tickets and %v are still left\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend")

	bookings := []string{}

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address:")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		

	

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - uint(userTickets)

			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("the first names bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("our conference is booked out. come back next year")
				break
			}

		} else {
			fmt.Printf("we only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			continue

		}

	}

}
