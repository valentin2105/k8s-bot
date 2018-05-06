package main

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/tbruyelle/hipchat-go/hipchat"
)

// GetLastHipchatMsg - Return last hipchat string msg
func GetLastHipchatMsg(baseurl string, token string, room string) string {
	var lastmsg string
	c := hipchat.NewClient(token)
	if baseurl != "" {
		url, err := url.Parse(baseurl)
		if err != nil {
			log.Printf("Error parsing hipchat base url: %s\n", err)
		}
		c.BaseURL = url
	}
	hist, resp, err := c.Room.History(room, &hipchat.HistoryOptions{})
	if err != nil {
		fmt.Printf("Error during room history req %q, Return : %+v \n Wait for 10 seconds... \n", err, resp)
		time.Sleep(10 * time.Second)
		lastmsg = "null"
	} else {
		last := hist.Items[len(hist.Items)-1]
		lastmsg = last.Message
	}
	return lastmsg

}

// HipchatNotify - Notify hipchat
func HipchatNotify(message string) {
	c := hipchat.NewClient(*token)
	notifRq := &hipchat.NotificationRequest{Message: message, Color: color, From: botName, MessageFormat: "text"}
	err, _ := c.Room.Notification(*room, notifRq)
	if err != nil {
	}
}
