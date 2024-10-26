package main

import (
	"fmt"
	"strings"

	"github.com/gunzette/fetchette/modules"
	"github.com/gunzette/fetchette/termcolor"
)

func main() {
	logo := termcolor.GetColoredLogoString("archColors.json")

	for i, line := range strings.Split(logo, "\n") {
		if line == "" {
			continue
		}

		fmt.Println(line, "a", i)
	}

	fmt.Println(modules.GetOS())
	fmt.Println(modules.GetUserAtHost())
	fmt.Println(modules.GetKernel())
}
