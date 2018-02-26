package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

const (
	kubeWord    = "!k"
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
			cmd := CheckBeforeExec(words, lastmsg)
			if cmd != "null" {
				fmt.Printf("----> Command executed : %+v\n", cmd)
				// Let's launch the kubectl cmd.
				args := strings.Split(cmd, " ")
				out, err := exec.Command(args[0], args[1:]...).Output()
				result := fmt.Sprintf("/code %s", out)
				cl := strings.Replace(result, "\n\n", "\n", -1)
				if err == nil {
					HipchatNotify(cl)
					fmt.Println("Hipchat message sended...")
				} else {
					fmt.Printf("Error during kubectl cmd : %q \n", err)
				}
				BasicAnswers(words)
			}
		} else {
			fmt.Printf("Error, Provider unavailable %s \n", provider)
		}
		// Slack
		if *provider == "Slack" {
			// Nothing for now...
		}
		time.Sleep(watchSecond * time.Second)
	}
}
