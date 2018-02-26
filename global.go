package main

import (
	"fmt"
	"strings"
)

// BasicAnswers - ...
func BasicAnswers(firstword []string) {
	// Add some funny questions
	if firstword[0] == "Bonjour" || firstword[0] == "Hello" {
		HipchatNotify("Bienvenue patron, que puis-je faire pour vous ?")
	}
	if firstword[0] == "help" || firstword[0] == "aide" || firstword[0] == "?" {
		HipchatNotify("Pour interagir avec moi : '!k namespace verb ressource' (!k default get pod) ...")
	}
	if firstword[0] == "Qui" {
		HipchatNotify("Vous, bien sûr !")
	}
	if firstword[0] == "Merci" {
		HipchatNotify("Avec plaisir patron.")
	}
}

// CheckBeforeExec - Check stuffs before exec.
func CheckBeforeExec(words []string, lastmsg string) string {
	if words[0] == kubeWord {
		cmd := strings.Replace(lastmsg, kubeWord, "kubectl -n", -1)
		// If it contain "all" namespace
		if words[1] == "all" {
			cmd = cmd + " --all-namespaces"
		}
		// If command is too short
		if len(words) <= 3 {
			fmt.Printf("Error, command unavailable %+v \n", cmd)
			HipchatNotify("Error, command incomplete")
			cmd = "null"
		}
		// Match TRUSTED words (get, scale ...)
		if StringInSlice(words[2], trustedVerbs) {
			if words[2] == "logs" && StringInSlice("-f", words) {
				fmt.Printf("Error, command unavailable %+v \n", cmd)
				HipchatNotify("Error, command Forbidden (logs -f)")
				cmd = "null"
			}
			if words[2] == "exec" && StringInSlice("-it", words) {
				fmt.Printf("Error, command unavailable %+v \n", cmd)
				HipchatNotify("Error, command Forbidden (exec -it)")
				cmd = "null"
			}
		}
		return cmd
	} else {
		cmd = "null"
		return cmd
	}
}

// StringInSlice - check string in slice
func StringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}
