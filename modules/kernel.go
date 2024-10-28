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

func GetKernel() string {
	return "{01}Kernel{02}: " + getKernelLinux() + "{-1}"
}
