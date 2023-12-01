package main

import (
	"fmt"
	"time"
)

func addEmailsToQueue(emails []string, ch chan string) {
	for _, email := range emails {
		ch <- email
	}
}

// TEST SUITE - Don't Touch Below This Line

func sendEmails(batchSize int, ch chan string) {
	for i := 0; i < batchSize; i++ {
		email := <-ch
		fmt.Println("Sending email:", email)
	}
}

func test(emails ...string) {
	fmt.Printf("Adding %v emails to queue...\n", len(emails))
	ch := make(chan string)
	go addEmailsToQueue(emails, ch)
	fmt.Println("Sending emails...")
	go sendEmails(len(emails), ch)
	fmt.Println("==========================================")
}

func main() {
	test("Hello John, tell Kathy I said hi", "Whazzup bruther")
	test("I find that hard to believe.", "When? I don't know if I can", "What time are you thinking?")
	test("She says hi!", "Yeah its tomorrow. So we're good.", "Cool see you then!", "Bye!")
	time.Sleep(time.Millisecond * 10) // ensure the last print statement happens
}
