package modules

import (
	"log"
	"os/exec"
	"runtime"
	"strings"
)

type OS int8

const (
	Windows OS = iota
	Linux
	Other
)

func getOSLinux() string {
	out, err := exec.Command("lsb_release", "-d").Output()
	if err != nil {
		log.Fatal("Error while getting os information: ", err)
	}
	res := strings.Split(string(out), "\t")
	return strings.Trim(res[1], "\n")
}

func GetOSString(curOS OS) string {
	switch curOS {
	case Windows:
		return "{01}OS{02}: Windows{-1}"
	case Linux:
		return "{01}OS{02}: " + getOSLinux() + "{-1}"
	default:
		return "{01}OS{02}: Unknown{-1}"
	}
}

func GetOS() OS {
	switch runtime.GOOS {
	case "windows":
		return Windows
	case "linux":
		return Linux
	default:
		return Other
	}
}
