package main;

import (
	"fmt"
);

func main(){
	var conferenceName string = "Go conference";
	const conferenceTickets  = 50;
	var remainingTickets int = 50;

	var firstName string;
	var lastName string;
	var email string;
	var userTickets int;

	fmt.Printf("Welcome to %v booking application\n", conferenceName);
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets);
	fmt.Println("Get your tickets here to attend");


	fmt.Println("Enter your firstname");
	fmt.Scan(&firstName);

	fmt.Println("Enter your lastname");
	fmt.Scan(&lastName);

	fmt.Println("Enter your email");
	fmt.Scan(&email);

	fmt.Println("How many tickets would you want?");
	fmt.Scan(&userTickets);

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v \n", firstName, lastName, userTickets,email )
}


