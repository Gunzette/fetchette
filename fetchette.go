package main

import (
	"fmt"

	"github.com/gunzette/fetchette/modules"
	"github.com/gunzette/fetchette/termcolor"
)

func main() {
	fmt.Println(termcolor.GetColoredLogoString("lesbianColors.json"))
	fmt.Println(modules.GetOS())
	fmt.Println(modules.GetUserAtHost())
	fmt.Println(modules.GetKernel())
}
