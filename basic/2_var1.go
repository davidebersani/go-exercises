package main

import "fmt"

func main() {

	// Declare variables.
	// Default to 0, false and empty string
	var smsSendingLimit uint64
	var costPerSms float64
	var hasPermission bool
	var username string

	fmt.Printf(
		"%v %f %v %q\n",
		smsSendingLimit,
		costPerSms,
		hasPermission,
		username,
	)

	// Declare variables and let Go infers the type
	// Type can't change! Go infers types

	congrats := "Happy birthday!"
	fmt.Println(congrats)

	// Declare a float

	penniesPerText := 2.0
	fmt.Printf("The type of penniesPerText is %T\n", penniesPerText)

	// Multiple declaration
	averageOpenRate, displayMessage := .23, "is the average open rate of your messages"
	fmt.Println(averageOpenRate, displayMessage)

	// Convert types
	accountAge := 2.6
	accountAgeInt := int(accountAge)

	fmt.Println("Your account has existed for", accountAgeInt, "years")

	// Constants
	const premiumPlanName = "Premium Plan"
	const basicPlanName = "Basic Plan"

	fmt.Println("plan:", premiumPlanName)
	fmt.Println("plan:", basicPlanName)

	// Il valore delle costanti può essere anche calcolato, ma deve essere possibile calcolarlo a tempo di compilazione
	const secondsInMinute = 60
	const minutesInHour = 60
	const secondsInHour = secondsInMinute * minutesInHour

	fmt.Println("number of seconsa in an hour:", secondsInHour)

	// Formatting
	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent", "Davide", 30.86)
	fmt.Println(msg)

	// Conditionals
	messageLen := 10
	maxMessageLen := 20
	fmt.Println("Trying to send a message of length:", messageLen)

	if messageLen <= maxMessageLen {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}

	// Syntactic sugar: if INITIAL_STATEMENT; CONDITION { }
	// Le variabili definite nel primo statement hanno vita solo dentro l'if
}
