package modules

import (
	"log"
	"os/exec"
	"strings"
)

func getHostname() string {
	out, err := exec.Command("hostname").Output()
	if err != nil {
		log.Fatal("Error while getting hostname: ", err)
	}
	res := strings.Trim(string(out), "\n")
	return res
}

func getUsername() string {
	out, err := exec.Command("whoami").Output()
	if err != nil {
		log.Fatal("Error while getting username: ", err)
	}
	res := strings.Trim(string(out), "\n")
	return res
}

func GetUserAtHost(curOS OS) string {
	switch curOS {
	default:
		return "{00}" + getUsername() + "{02}@{00}" + getHostname() + "{-1}"
	}
}
