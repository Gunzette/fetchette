package main

import (
	"fmt"

	"github.com/gunzette/fetchette/termcolor"
)

func main() {
	fmt.Println(termcolor.BgColor("   ", &[3]int{20, 255, 2}))
	fmt.Println(termcolor.FullColor("Haiii", &[3]int{0, 255, 0}, &[3]int{255, 0, 0}))
}
