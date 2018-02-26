package main

import (
	"fmt"
	"os/exec"
	"strings"
)

// CheckBeforeExec - Check stuffs before exec.
func CheckBeforeExec(words []string, lastmsg string) string {
	cmd := "null"
	if words[0] == KubeWord && len(words) >= 3 {
		cmd = strings.Replace(lastmsg, KubeWord, "/usr/local/bin/kubectl -n", -1)

		// If it contain "all" namespace
		if words[1] == "all" {
			cmd = cmd + " --all-namespaces"
		}

		if !StringInSlice(words[2], trustedVerbs) {
			fmt.Printf("-> Error, command unavailable %+v \n", cmd)
			HipchatNotify("Error, command Forbidden")
			cmd = "null"
		}
		// Match TRUSTED words (get, scale ...)
		if words[2] == "logs" && StringInSlice("-f", words) {
			fmt.Printf("-> Error, command unavailable %+v \n", cmd)
			HipchatNotify("Error, command Forbidden (logs -f)")
			cmd = "null"
		}
		if words[2] == "exec" && StringInSlice("-it", words) {
			fmt.Printf("-> Error, command unavailable %+v \n", cmd)
			HipchatNotify("Error, command Forbidden (exec -it)")
			cmd = "null"
		}
	}
	return cmd
}

// ExecKubectl - Launch and format kubectl cmd.
func ExecKubectl(cmd string) string {
	cl := "null"
	args := strings.Split(cmd, " ")
	out, err := exec.Command(args[0], args[1:]...).Output()
	if err == nil {
		result := fmt.Sprintf("/code %s", out)
		cl = strings.Replace(result, "\n\n", "\n", -1)
	}
	return cl
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
