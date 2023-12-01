package main

import "fmt"

type messageToSend struct {
	Message   string
	Sender    user
	Recipient user
}

type user struct {
	Name   string
	Number int
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.Sender.Name == "" || mToSend.Sender.Number == 0 || mToSend.Recipient.Name == "" || mToSend.Recipient.Number == 0 {
		return false
	}

	return true
}

func test(m messageToSend) {
	if canSendMessage(m) {
		fmt.Printf("Sending message: '%s' from %s (%v) to %s (%v)\n", m.Message, m.Sender.Name, m.Sender.Number, m.Recipient.Name, m.Recipient.Number)
		fmt.Println("============================================")
	}
}

func main() {
	test(messageToSend{
		Message: "Hello!",
		Sender: user{
			Name:   "Davide",
			Number: 125615,
		},
		Recipient: user{
			Name:   "Franco",
			Number: 156486,
		},
	})
	test(messageToSend{
		Message: "This should not be sent",
		Sender: user{
			Name:   "",
			Number: 0,
		},
		Recipient: user{
			Name:   "Franco",
			Number: 156486,
		},
	})

}
