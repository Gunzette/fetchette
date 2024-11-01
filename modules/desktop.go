package modules

import (
	"os"
)

func getDesktopLinux() string {
	out := os.Getenv("XDG_CURRENT_DESKTOP")
	return out
}

func GetDesktop(curOS OS) string {
	switch curOS {
	// Technically wrong, because you can use window managers on windows but i do not know how to check
	case Windows:
		return "{01}Desktop{02}: Windows Shell{-1}"
	case Linux:
		return "{01}Desktop{02}: " + getDesktopLinux() + "{-1}"
	default:
		return "{01}Desktop{02}: Unknown{-1}"
	}
}
