package modules

import (
	"os"
)

func getDesktopLinux() string {
	out := os.Getenv("XDG_CURRENT_DESKTOP")
	return out
}

func GetDesktop() string {
	return "{01}Desktop{02}: " + getDesktopLinux() + "{-1}"
}
