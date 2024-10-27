package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/gunzette/fetchette/modules"
	"github.com/gunzette/fetchette/termcolor"
)

const LOGOGAP int = 2

func displayFetch(colorsFile string, moduleOutputs []string) {
	// Get logo string
	logo := termcolor.GetColoredLogoString(colorsFile)

	outputLen := len(moduleOutputs)

	// Add to each line of logo
	for i, line := range strings.Split(logo, "\n") {
		if line == "" {
			continue
		}

		if i < outputLen {
			fmt.Println(line, strings.Repeat(" ", LOGOGAP), moduleOutputs[i])
		} else {
			fmt.Println(line)
		}
	}
}

func main() {
	// Check arguments
	if len(os.Args) == 1 {
		log.Fatal("Not enough arguments")
	}

	userAtHost := modules.GetUserAtHost()

	displayFetch(os.Args[1]+"Colors.json", []string{userAtHost, strings.Repeat("-", utf8.RuneCountInString(userAtHost)), modules.GetOS(), modules.GetKernel()})
}
