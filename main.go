package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/tbruyelle/hipchat-go/hipchat"
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
		if *provider == "Hipchat" {
			c := hipchat.NewClient(*token)
			hist, resp, err := c.Room.History(*room, &hipchat.HistoryOptions{})
			if err != nil {
				fmt.Printf("Error during room history req %q, Return : %+v \n Wait for 10 seconds... \n", err, resp)
				time.Sleep(10 * time.Second)
				goto start
			}
			last := hist.Items[len(hist.Items)-1]
			lastmsg := last.Message
			words := strings.Fields(lastmsg)
			// If the last message contain the trigger "kubeWord"
			if words[0] == kubeWord {
				cmd = strings.Replace(lastmsg, kubeWord, "kubectl -n", -1)
				// If it contain "all" namespace
				if words[1] == "all" {
					cmd = cmd + " --all-namespaces"
				}
				// If command is too short
				if len(words) <= 3 {
					fmt.Print("Error, command unavailable %+v \n", cmd)
					HipchatNotify("Error, command incomplete")
				}
				// Match TRUSTED words (get, scale ...)
				if StringInSlice(words[2], trustedVerbs) {
					if words[2] == "logs" && StringInSlice("-f", words) {
						fmt.Print("Error, command unavailable %+v \n", cmd)
						HipchatNotify("Error, command Forbidden (logs -f)")
					}
					if words[2] == "exec" && StringInSlice("-it", words) {
						fmt.Print("Error, command unavailable %+v \n", cmd)
						HipchatNotify("Error, command Forbidden (exec -it)")
					}
					// Let's launch the kubectl cmd.
					fmt.Printf("----> Command executed : %+v\n", cmd)
					args := strings.Split(cmd, " ")
					out, err := exec.Command(args[0], args[1:]...).Output()
					result := fmt.Sprintf("/code %s", out)
					cl := strings.Replace(result, "\n\n", "\n", -1)
					if err != nil {
						fmt.Printf("Error during kubectl cmd : %q \n", err)
					}
					if err == nil {
						HipchatNotify(cl)
						fmt.Println("Hipchat Notified...")
					}
				} else {
					fmt.Print("Error, command unavailable %+v \n", cmd)
					HipchatNotify("Error, command unavailable...")
				}
			} else {
				BasicAnswers(words[0])
			}
		} else {
			fmt.Print("Error, Provider unavailable %s \n", provider)
		}
		// Slack
		if *provider == "Slack" {
		}
		time.Sleep(watchSecond * time.Second)
	}
}
