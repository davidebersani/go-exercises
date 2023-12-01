package main

import (
	"fmt"
	"time"
)

func sendMessage(msg message) {
	fmt.Println(msg.getMessage())
}

type message interface {
	getMessage() string
}

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime)
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf("Report pronto: %s", sr.reportName)
}

func test(m message) {
	sendMessage(m)
	fmt.Println("============================================")
}

func main() {
	test(sendingReport{
		reportName:    "Test",
		numberOfSends: 231,
	})

	test(birthdayMessage{
		recipientName: "Davide",
		birthdayTime:  time.Date(2001, 5, 6, 6, 7, 0, 0, time.UTC),
	})

}
