package main

import (
	"fmt"

	"github.com/tbruyelle/hipchat-go/hipchat"
)

// Add some funny questions
func BasicAnswers(firstword string) {
	if firstword == "Hello" {
		HipchatNotify("Hello boss, How I can help you ?")
	}
	if firstword == "help" || firstword == "?" {
		HipchatNotify("To request myself -> '!k namespace verb ressource' (!k default get pod) ...")
	}
	if firstword == "Who" {
		HipchatNotify("You, are !")
	}
	if firstword == "thanks" || firstword == "Thanks" {
		HipchatNotify("With pleasures, boss.")
	}
}

// Notify Hipchat
func HipchatNotify(message string) {
	c := hipchat.NewClient(*token)
	notifRq := &hipchat.NotificationRequest{Message: message, Color: color, From: botName, MessageFormat: "text"}
	err, _ := c.Room.Notification(*room, notifRq)
	if err != nil {
		fmt.Println("Error to Notify Hipchat...")
	}
}

// IsStringinSlice ?
func StringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}
