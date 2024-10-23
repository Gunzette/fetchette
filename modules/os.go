package modules

import (
	"log"
	"os/exec"
	"strings"
)

func getOSLinux() string {
	out, err := exec.Command("lsb_release", "-d").Output()
	if err != nil {
		log.Fatal("Error while getting os information: ", err)
	}
	res := strings.Split(string(out), "\t")
	return strings.Trim(res[1], "\n")
}

func GetOS() string {
	return "OS: " + getOSLinux()
}
