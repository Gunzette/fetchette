package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gunzette/fetchette/modules"
	"github.com/gunzette/fetchette/termcolor"
)

func main() {
	// Check arguments
	if len(os.Args) == 1 {
		log.Fatal("Not enough arguments")
	}

	// Get logo string
	logo := termcolor.GetColoredLogoString(os.Args[1] + "Colors.json")

	// Add to each line of logo
	for i, line := range strings.Split(logo, "\n") {
		if line == "" {
			continue
		}

		fmt.Println(line, "a", i)
	}

	// test modules
	fmt.Println(modules.GetOS())
	fmt.Println(modules.GetUserAtHost())
	fmt.Println(modules.GetKernel())
}
