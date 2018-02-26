package main

import (
	"fmt"
	"time"

	"github.com/tbruyelle/hipchat-go/hipchat"
)

// GetLastHipchatMsg - Return last hipchat string msg
func GetLastHipchatMsg(token string, room string) string {
	c := hipchat.NewClient(token)
	hist, resp, err := c.Room.History(room, &hipchat.HistoryOptions{})
	last := hist.Items[len(hist.Items)-1]
	lastmsg := last.Message
	if err != nil {
		fmt.Printf("Error during room history req %q, Return : %+v \n Wait for 10 seconds... \n", err, resp)
		time.Sleep(10 * time.Second)
		lastmsg = "null"
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
