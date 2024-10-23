package main

import (
	"fmt"

	"github.com/gunzette/fetchette/termcolor"
)

func main() {
	fmt.Println(termcolor.GetColoredLogoString("lesbianColors.json"))
}
