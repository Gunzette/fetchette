package modules

import (
	"log"
	"os/exec"
	"strings"
)

func getHostnameLinux() string {
	out, err := exec.Command("hostname").Output()
	if err != nil {
		log.Fatal("Error while getting hostname: ", err)
	}
	res := strings.Trim(string(out), "\n")
	return res
}

func getUsernameLinux() string {
	out, err := exec.Command("whoami").Output()
	if err != nil {
		log.Fatal("Error while getting username: ", err)
	}
	res := strings.Trim(string(out), "\n")
	return res
}

func GetUserAtHost() string {
	return "{00}" + getUsernameLinux() + "{02}@{00}" + getHostnameLinux() + "{-1}"
}
