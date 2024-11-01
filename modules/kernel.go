package modules

import (
	"log"
	"os/exec"
	"strings"
)

func getKernelLinux() string {
	out, err := exec.Command("uname", "-r").Output()
	if err != nil {
		log.Fatal("Error while getting kernel release", err)
	}
	res := strings.Trim(string(out), "\n")
	return res
}

func getKernelWindows() string {
	out, err := exec.Command("cmd", "-c", "ver").Output()
	if err != nil {
		log.Fatal("Error while getting windows release", err)
	}
	res := strings.Split(string(out), "\r\n")
	return res[0]
}

func GetKernel(curOS OS) string {
	switch curOS {
	case Windows:
		return "{01}Kernel{02}: " + getKernelWindows() + "{-1}"
	case Linux:
		return "{01}Kernel{02}: " + getKernelLinux() + "{-1}"
	default:
		return "{01}Kernel{02}: Unknown{-1}"
	}
}
