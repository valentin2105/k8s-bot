package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

const (
	// KubeWord - The Word that trigger the bot.
	KubeWord    = "!k"
	color       = "green"
	botName     = "k8s-bot"
	watchSecond = 3
)

var (
	trustedVerbs = []string{"get", "scale", "exec", "describe", "label", "annotate", "version", "logs", "rollout"}
	provider     = flag.String("provider", "Hipchat", "The Provider (Hipchat by default)")
	token        = flag.String("token", "", "The AuthToken")
	room         = flag.String("room", "", "The Room ID")
	cmd          string
)

func main() {
	flag.Parse()
	if *token == "" || *room == "" || *provider == "" {
		flag.PrintDefaults()
		return
	}

	for {
	start:
		// Hipchat
		if *provider == "Hipchat" {
			lastmsg := GetLastHipchatMsg(*token, *room)
			if lastmsg == "null" {
				goto start
			}
			words := strings.Fields(lastmsg)
			// Launch Checks
			cmd = CheckBeforeExec(words, lastmsg)
			if cmd != "null" {
				fmt.Printf("----> Command executed : %+v\n", cmd)
				// Let's launch the kubectl cmd.
				cl := ExecKubectl(cmd)
				if cl != "null" {
					HipchatNotify(cl)
					fmt.Println("--> Hipchat message send.")
				} else {
					fmt.Printf("Error during kubectl cmd. \n")
				}
			}
		} else {
			fmt.Printf("Error, Provider unavailable (Hipchat or Slack) \n")
		}
		// Slack
		if *provider == "Slack" {
			// Nothing for now...
		}
		time.Sleep(watchSecond * time.Second)
	}
}
