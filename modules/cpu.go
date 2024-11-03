package modules

import (
	"log"
	"os/exec"
	"strings"
)

func getCPULinux() string {
	out, err := exec.Command("lscpu", "-e=MODELNAME").Output()
	if err != nil {
		log.Fatal("Error while getting CPU information: ", err)
	}
	res := strings.Split(string(out), "\n")
	return res[2]
}

func getCPUWindows() string {
	out, err := exec.Command("wmic", "cpu", "get", "name").Output()
	if err != nil {
		log.Fatal("Error while getting CPU information: ", err)
	}
	res := strings.Split(string(out), "\r\n")
	return res[1]
}

func GetCPU(curOS OS) string {
	switch curOS {
	case Windows:
		return "{01}CPU{02}: " + getCPUWindows() + "{-1}"
	case Linux:
		return "{01}CPU{02}: " + getCPULinux() + "{-1}"
	default:
		return "{01}CPU{02}: Unknown{-1}"
	}
}
